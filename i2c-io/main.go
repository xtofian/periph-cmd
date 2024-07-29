// Copyright 2016 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// i2c-io communicates to an I²C device.
package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/conn/v3/i2c/i2ctest"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/host/v3"
)

// adapted from https://gist.github.com/chmike/05da938833328a9a94e02506922f2e7b
func xxd(b []byte) string {
	var a [16]byte
	s := ""
	n := (len(b) + 15) &^ 15
	for i := 0; i < n; i++ {
		if i%16 == 0 {
			s += fmt.Sprintf("%4d", i)
		}
		if i%8 == 0 {
			s += " "
		}
		if i < len(b) {
			s += fmt.Sprintf(" %02X", b[i])
		} else {
			s += "   "
		}
		if i >= len(b) {
			a[i%16] = ' '
		} else if b[i] < 32 || b[i] > 126 {
			a[i%16] = '.'
		} else {
			a[i%16] = b[i]
		}
		if i%16 == 15 {
			s += fmt.Sprintf("  %s\n", string(a[:]))
		}
	}
	return s
}

func mainImpl() error {
	addr := flag.Int("a", -1, "I²C device address to query")
	busName := flag.String("b", "", "I²C bus to use")
	verbose := flag.Bool("v", false, "verbose mode")
	// TODO(maruel): This is not generic enough.
	write := flag.Bool("w", false, "write instead of reading")
	reg := flag.Int("r", -1, "register to address")
	var hz physic.Frequency
	flag.Var(&hz, "hz", "I²C bus speed (may require root)")
	l := flag.Int("l", 1, "length of data to read; ignored if -w is specified")
	flag.Parse()
	if !*verbose {
		log.SetOutput(ioutil.Discard)
	}
	log.SetFlags(log.Lmicroseconds)
	if !*write && flag.NArg() != 0 {
		return errors.New("unexpected argument, try -help")
	}

	if *addr < 0 || *addr >= 1<<9 {
		return fmt.Errorf("-a is required and must be between 0 and %d", 1<<9-1)
	}
	if *reg < -1 || *reg > 255 {
		return errors.New("-r must be between -1 and 255")
	}
	if *l <= 0 || *l > 255 {
		return errors.New("-l must be between 1 and 255")
	}

	if _, err := host.Init(); err != nil {
		return err
	}

	var buf []byte
	if *write {
		if flag.NArg() == 0 {
			return errors.New("specify data to write as a list of hex encoded bytes")
		}
		if *reg >= 0 {
			buf = make([]byte, 1, flag.NArg()+1)
			buf[0] = byte(*reg)
		} else {
			buf = make([]byte, 0, flag.NArg())
		}
		for _, a := range flag.Args() {
			b, err := strconv.ParseUint(a, 16, 8)
			if err != nil {
				return err
			}
			buf = append(buf, byte(b))
		}
	} else {
		if flag.NArg() != 0 {
			return errors.New("do not specify bytes when reading")
		}
		buf = make([]byte, *l)
	}

	bus, err := i2creg.Open(*busName)
	if err != nil {
		return err
	}
	defer bus.Close()

	if hz != 0 {
		if err = bus.SetSpeed(hz); err != nil {
			return err
		}
	}
	if *verbose {
		if p, ok := bus.(i2c.Pins); ok {
			log.Printf("Using pins SCL: %s  SDA: %s", p.SCL(), p.SDA())
		}
	}
	i2cRecorder := i2ctest.Record{Bus: bus}

	d := i2c.Dev{Bus: &i2cRecorder, Addr: uint16(*addr)}
	if *write {
		if *verbose {
			log.Printf("Writing: %#v", buf)
		}
		_, err = d.Write(buf)
	} else {
		req := []byte{}
		if *reg >= 0 {
			req = []byte{byte(*reg)}
		}
		if err = d.Tx(req, buf); err != nil {
			return err
		}
		for i, b := range buf {
			if i != 0 {
				if _, err = fmt.Print(", "); err != nil {
					break
				}
			}
			if _, err = fmt.Printf("0x%02X", b); err != nil {
				break
			}
		}
		_, err = fmt.Print("\n")
	}

	if *verbose {
		log.Printf("recorded ops:\n%#v", i2cRecorder.Ops)

		for _, op := range i2cRecorder.Ops {
			log.Printf("  W:\n%s", xxd(op.W))
			log.Printf("  R:\n%s", xxd(op.R))
		}
	}
	return err
}

func main() {
	if err := mainImpl(); err != nil {
		fmt.Fprintf(os.Stderr, "i2c-io: %s.\n", err)
		os.Exit(1)
	}
}
