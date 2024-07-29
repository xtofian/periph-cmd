// Copyright 2016 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// ezo communicates to an EZO device via I²C.
package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/devices/v3/ezo"
	"periph.io/x/host/v3"
)

func mainImpl() error {
	busName := flag.String("b", "I2C1", "I²C bus to use")
	verbose := flag.Bool("v", false, "verbose mode")

	flag.Parse()

	if flag.NArg() < 2 {
		flag.Usage()
		return errors.New("need command")
	}

	ezoDevice := flag.Arg(0)
	cmd := flag.Arg(1)

	var addr uint16
	switch ezoDevice {
	case "ph":
		addr = ezo.EzoPhI2CAddr
	case "orp":
		addr = ezo.EzoOrpI2CAddr
	default:
		flag.Usage()
		return fmt.Errorf("invalid device type: %s", ezoDevice)
	}

	if *verbose {
		log.Printf("Open i2c bus \"%s for", *busName)
	}

	if _, err := host.Init(); err != nil {
		return err
	}

	i2c, err := i2creg.Open(*busName)
	if err != nil {
		return err
	}
	defer i2c.Close()

	ezoDev, err := ezo.NewEzo(i2c, addr)
	if err != nil {
		return err
	}

	switch cmd {
	case "rawcmd":
		if flag.NArg() != 3 {
			return errors.New("rawcmd: missing argument")
		}
		resp, err := ezoDev.Command(flag.Arg(2), 900*time.Millisecond)
		if err != nil {
			return err
		}

		fmt.Println(resp)

	default:
		return fmt.Errorf("unrecognized command: %s", cmd)

	}

	return nil
}

func main() {
	if err := mainImpl(); err != nil {
		fmt.Fprintf(os.Stderr, "ezo: %s.\n", err)
		os.Exit(1)
	}
}
