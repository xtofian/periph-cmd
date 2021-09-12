// Code generated by "go run internal/static_gen.go -o static_prod.go"; DO NOT EDIT.

//go:build !debug
// +build !debug

package main

const cacheControlContent = "Cache-Control:public,max-age=300"

func getContent(path string) []byte {
	return staticContent[path]
}

var staticContent = map[string][]byte{
	"static/index.html":  []byte("<!doctype html><meta charset=utf-8><meta name=viewport content=\"width=device-width,initial-scale=1\"><meta name=apple-mobile-web-app-capable content=yes><meta name=apple-mobile-web-app-status-bar-style content=blue><meta name=apple-mobile-web-app-capable content=yes><meta name=mobile-web-app-capable content=yes><meta name=description content=\"Periph web UI\"><meta name=author content=\"Periph Authors\"><title>periph-web</title><style>*{font-family:sans-serif;font-size:14px}h1{font-size:24px}h2{font-size:20px}h3{font-size:16px}h1,h2,h3{margin-bottom:.2em;margin-top:.2em}.err{background:#f44;border:1px solid #888;border-radius:10px;padding:10px;display:none}@media only screen and (max-width:500px){*{font-size:12px}}</style><script>\"use strict\";function log(v){}\nclass EventSource{constructor(){this._triggers={};}\naddEventListener(type,listener,options){if(!this._triggers[type]){this._triggers[type]=[];}\nlet opt=options||{};let v={capture:opt.capture,listener:listener,once:opt.once,passive:opt.passive,};this._triggers[type].push(v);}\nremoveEventListener(type,listener,options){if(!this._triggers[type]){return;}\nlet l=this._triggers[type].slice();let opt=options||{};for(let i=l.length;i>0;i--){let v=l[i-1];if(v.callback===callback&&v.capture===opt.capture&&v.passive===opt.passive){this._triggers[type].pop(i);}}}\ndispatchEvent(type,params){log(\"dispatchEvent(\"+type+\", \"+params+\")\");let l=this._triggers[type];if(!l){return;}\nlet rm=[];for(let i=0;i<l.length;i++){let opt=l[i];opt.listener.call(params);if(opt.once){rm.push(opt);}}\nfor(let i=0;i<rm.length;i++){for(let j=0;j<l.length;l++){if(l[j]===rm[i]){l.pop(j);break;}}}}}\nfunction postJSON(url,data,callback){function checkStatus(res){if(res.status==401){throw new Error(\"Please refresh the page\");}\nif(res.status>=200&&res.status<300){return res.json();}\nthrow new Error(res.statusText);}\nfunction onError(url,err){console.log(err);alertError(url+\": \"+err.toString());}\nlet hdr={body:JSON.stringify(data),credentials:\"same-origin\",headers:{\"Content-Type\":\"application/json; charset=utf-8\"},method:\"POST\",};fetch(url,hdr).then(checkStatus).then(callback).catch(err=>onError(url,err));}\nfunction alertError(errText){let e=document.getElementById(\"err\");if(e.innerText){e.innerText=e.innerText+\"\\n\";}\ne.innerText=e.innerText+errText+\"\\n\";e.style.display=\"block\";}\nclass Pin{constructor(name,number,func,gpio){this.name=name;this.number=number;this._func=func;this.gpio=gpio;}\nget func(){if(this.gpio){return this.gpio.func;}\nreturn this._func;}}\nclass GPIO{constructor(name,number,func){this.name=name;this.number=number;this.func=func;this._value=null;}\nget value(){return this._value;}\nonFuncUpdate(f,v){if(this.func===this._makeFunc(f,v)){return false;}\nthis._value=v;this.func=this._makeFunc(f,v);return true;}\nonValueRead(v){if(this._value===v){return false;}\nif(this.type===\"out\"){return false;}\nthis._value=v;this.func=this._makeFunc(this.type,v);return true;}\nget type(){if(this.func.startsWith(\"Out/\")){return \"out\";}\nif(this.func.startsWith(\"In/\")){return \"in\";}\nreturn this.func;}\n_makeFunc(t,v){if(t==\"in\"){if(v===true){return \"In/High\";}else if(v===false){return \"In/Low\";}\nreturn \"In/Ind\";}else if(t==\"out\"){if(v===true){return \"Out/High\";}else if(v===false){return \"Out/Low\";}\nreturn \"Out/Ind\";}\nreturn t;}}\nclass Header{constructor(name,pins){this.name=name;this.pins=pins;}}\nvar Controller=new class{constructor(){this.eventGPIO=new EventSource();this.eventDone=new EventSource();this.gpios={};this.headers={};this._polling={};this._pollingID=null;this._pollingRate=100;document.addEventListener(\"DOMContentLoaded\",()=>{this._fetchGPIO();this._fetchHeader();},{once:true});}\nsetGPIOIn(gpio){log(\"setGPIOIn(\"+gpio.name+\")\");let params={Name:gpio.name,Pull:\"\",Edge:\"\",};gpio.onFuncUpdate(\"in\",null);this.eventGPIO.dispatchEvent(gpio.name,\"in\");postJSON(\"/api/periph/v1/gpio/in\",[params],res=>{if(res[0]){alertError(res[0]);return;}});}\nsetGPIOOut(gpio,l){log(\"setGPIOOut(\"+gpio.name+\", \"+l+\")\");let params={};params[gpio.name]=l;gpio.onFuncUpdate(\"out\",l);this.eventGPIO.dispatchEvent(gpio.name,\"out\");postJSON(\"/api/periph/v1/gpio/out\",params,res=>{if(res[0]){alertError(res[0]);return;}});}\n_fetchGPIO(){postJSON(\"/api/periph/v1/gpio/list\",{},res=>{log(\"/api/periph/v1/gpio/list\");for(let i=0;i<res.length;i++){let name=res[i].Name;let gpio=new GPIO(name,res[i].Number,res[i].Func);this.gpios[name]=gpio;this.eventGPIO.dispatchEvent(name);}\nthis.eventDone.dispatchEvent(\"gpio\");});}\n_fetchHeader(){postJSON(\"/api/periph/v1/header/list\",{},res=>{log(\"/api/periph/v1/header/list\");for(let key in res){let pins=[];for(let y=0;y<res[key].Pins.length;y++){let row=res[key].Pins[y];let items=[];for(let x=0;x<row.length;x++){let src=row[x];let p=new Pin(src.Name,src.Number,src.Func,this.gpios[src.Name]);if(!p.gpio){this.eventGPIO.addEventListener(p.name,()=>{p.gpio=this.gpios[p.name];this._autoPoll(p.gpio);this.eventGPIO.addEventListener(p.gpio.name,()=>{this._autoPoll(p.gpio);});},{once:true});}else{this._autoPoll(p.gpio);this.eventGPIO.addEventListener(p.gpio.name,()=>{this._autoPoll(p.gpio);});}\nitems[x]=p;}\npins[y]=items;}\nthis.headers[key]=new Header(key,pins);}\nthis.eventDone.dispatchEvent(\"header\");});}\n_autoPoll(gpio){if(gpio.type==\"in\"){log(\"Now polling \"+gpio.name);this._polling[gpio.name]=true;if(this._pollingID===null){this._pollingID=window.setTimeout(this._refreshGPIO.bind(this),this._pollingRate);}}else{log(\"Stop polling \"+gpio.name);delete this._polling[gpio.name];if(this._pollingID&&!this._polling.length){window.clearTimeout(this._pollingID);this._pollingID=null;}}}\n_refreshGPIO(){let pins=Object.keys(this._polling).sort();postJSON(\"/api/periph/v1/gpio/read\",pins,res=>{for(let i=0;i<pins.length;i++){let p=this.gpios[pins[i]];switch(res[i]){case 0:if(p.onValueRead(false)){this.eventGPIO.dispatchEvent(p.name,false);}\nbreak;case 1:if(p.onValueRead(true)){this.eventGPIO.dispatchEvent(p.name,true);}\nbreak;default:if(p.onValueRead(null)){this.eventGPIO.dispatchEvent(p.name,null);}\nbreak;}}\nthis._pollingID=setTimeout(this._refreshGPIO.bind(this),this._pollingRate);});}};function fetchI2C(){postJSON(\"/api/periph/v1/i2c/list\",{},res=>{let root=document.getElementById(\"section-i2c\");for(let i=0;i<res.length;i++){let e=root.appendChild(document.createElement(\"i2c-elem\"));e.setupI2C(res[i].Name,res[i].Number,res[i].Err,res[i].SCL,res[i].SDA);}});}\nfunction fetchSPI(){postJSON(\"/api/periph/v1/spi/list\",{},res=>{let root=document.getElementById(\"section-spi\");for(let i=0;i<res.length;i++){let e=root.appendChild(document.createElement(\"spi-elem\"));e.setupSPI(res[i].Name,res[i].Number,res[i].Err,res[i].CLK,res[i].MOSI,res[i].MISO,res[i].CS);}});}\nfunction fetchState(){postJSON(\"/api/periph/v1/server/state\",{},res=>{document.title=\"periph-web - \"+res.Hostname;let root=document.getElementById(\"section-drivers-loaded\");if(!res.State.Loaded.length){root.display=\"hidden\";}else{root.setupDrivers([\"Drivers loaded\"]);for(let i=0;i<res.State.Loaded.length;i++){root.appendRow([res.State.Loaded[i]]);}}\nroot=document.getElementById(\"section-drivers-skipped\");if(!res.State.Skipped.length){root.display=\"hidden\";}else{root.setupDrivers([\"Drivers skipped\",\"Reason\"]);for(let i=0;i<res.State.Skipped.length;i++){root.appendRow([res.State.Skipped[i].D,res.State.Skipped[i].Err]);}}\nroot=document.getElementById(\"section-drivers-failed\");if(!res.State.Failed.length){root.display=\"hidden\";}else{root.setupDrivers([\"Drivers failed\",\"Error\"]);for(let i=0;i<res.State.Failed.length;i++){root.appendRow([res.State.Failed[i].D,res.State.Failed[i].Err]);}}});}\nController.eventDone.addEventListener(\"header\",()=>{let root=document.getElementById(\"section-gpio\");Object.keys(Controller.headers).sort().forEach(key=>{root.appendChild(document.createElement(\"header-view\")).setupHeader(key);});},{once:true});document.addEventListener(\"DOMContentLoaded\",()=>{fetchI2C();fetchSPI();fetchState();},{once:true});class HTMLElementTemplate extends HTMLElement{constructor(template_name){super();let tmpl=document.querySelector(\"template#\"+template_name);this.attachShadow({mode:\"open\"}).appendChild(tmpl.content.cloneNode(true));}\nstatic get observedAttributes(){return[];}\nemitEvent(name,detail){this.dispatchEvent(new CustomEvent(name,{detail,bubbles:true}));}}</script><template id=template-data-table-elem><style>th{background-color:#4caf50;color:#fff}th,td{padding:.5rem;border-bottom:1px solid #ddd}tr:hover{background-color:#ccc}tr:nth-child(even):not(:hover){background:#f5f5f5}.inline{display:inline-block;margin-bottom:1rem;margin-right:2rem;vertical-align:top}</style><div class=inline><table><thead><tbody></table></div></template><script>\"use strict\";window.customElements.define(\"data-table-elem\",class extends HTMLElementTemplate{constructor(){super(\"template-data-table-elem\");}\nsetupTable(hdr){let root=this.shadowRoot.querySelector(\"thead\");for(let i=0;i<hdr.length;i++){root.appendChild(document.createElement(\"th\")).innerText=hdr[i];}}\nappendRow(line){let tr=this.shadowRoot.querySelector(\"tbody\").appendChild(document.createElement(\"tr\"));let items=[];for(let i=0;i<line.length;i++){let cell=tr.appendChild(document.createElement(\"td\"));if(line[i]instanceof Element){cell.appendChild(line[i]);items[i]=line[i];}else{cell.innerText=line[i];items[i]=cell;}}\nreturn items;}});</script><template id=template-checkout-elem><style>@keyframes popIn{0%{transform:scale(1,1)}25%{transform:scale(1.2,1)}50%{transform:scale(1.4,1)}100%{transform:scale(1,1)}}@keyframes popOut{0%{transform:scale(1,1)}25%{transform:scale(1.2,1)}50%{transform:scale(1.4,1)}100%{transform:scale(1,1)}}div{display:inline-block;height:20px;position:relative;vertical-align:bottom}input{bottom:0;cursor:pointer;display:block;height:0%;left:0;margin:0;opacity:0;position:absolute;right:0;top:0;width:0%}span{cursor:pointer;margin-left:.25em;padding-left:40px;user-select:none}span:before{background:#64646433;border-radius:20px;box-shadow:inset 0 0 5px #000c;content:\"\";display:inline-block;height:20px;left:0;position:absolute;transition:background .2s ease-out;width:40px}span:after{background-clip:padding-box;background:#fff;border-radius:20px;border:solid green 2px;content:\"\";display:block;font-weight:700;height:20px;left:-2px;position:absolute;text-align:center;top:-2px;transition:margin-left .1s ease-in-out;width:20px}input:checked+span:after{margin-left:20px}input:checked+span:before{transition:background .2s ease-in}input:not(:checked)+span:after{animation:popOut ease-in .3s normal}input:checked+span:after{animation:popIn ease-in .3s normal;background-clip:padding-box;margin-left:20px}input:checked+span:before{background:#20c997}input:disabled+span:before{box-shadow:0 0 0 #000}input:disabled+span{color:#adb5bd}input:disabled:checked+span:before{background:#adb5bd}input:indeterminate+span:after{margin-left:10px}input:focus+span:before{outline:solid #cce5ff 2px}</style><div><label><input type=checkbox><span><slot></slot></span></label></div></template><script>\"use strict\";window.customElements.define(\"checkout-elem\",class extends HTMLElementTemplate{constructor(){super(\"template-checkout-elem\");}\nconnectedCallback(){this.contentElem=this.shadowRoot.querySelector(\"span\");this.checkboxElem=this.shadowRoot.querySelector(\"input\");this.checkboxElem.addEventListener(\"click\",e=>{this.emitEvent(\"change\",{});},{passive:true});}\nget checked(){return this.checkboxElem.checked;}\nset checked(v){this.checkboxElem.checked=v;}\nget disabled(){return this.checkboxElem.disabled;}\nset disabled(v){this.checkboxElem.disabled=v;}\nget indeterminate(){return this.checkboxElem.indeterminate;}\nset indeterminate(v){this.checkboxElem.indeterminate=v;}\nget text(){return this.contentElem.innerText;}\nset text(v){this.contentElem.innerText=v;}});</script><template id=template-drivers-elem><style>.inline{display:inline-block}</style><div class=inline><data-table-elem></data-table-elem></div></template><script>\"use strict\";window.customElements.define(\"drivers-elem\",class extends HTMLElementTemplate{constructor(){super(\"template-drivers-elem\");}\nsetupDrivers(hdr){this.shadowRoot.querySelector(\"data-table-elem\").setupTable(hdr);}\nappendRow(row){this.shadowRoot.querySelector(\"data-table-elem\").appendRow(row);}});</script><template id=template-gpio-view><style>div{background:#ccc;border:1px solid #888;border-radius:10px;padding:10px}.gpio{background:#fcf}.controls{display:none}.box{border:1px solid #888;border-radius:6px;padding:3px}#func{display:none;background-color:#ccc;padding-bottom:3px;padding-right:3px;border-radius:3px}</style><div><span id=name>L</span>\n<span><span class=controls><span><checkout-elem id=io>I/O</checkout-elem>\n<checkout-elem id=level>Level</checkout-elem></span></span>\n<span id=func></span></span></div></template><script>\"use strict\";window.customElements.define(\"gpio-view\",class extends HTMLElementTemplate{constructor(){super(\"template-gpio-view\");}\nconnectedCallback(){this.funcElem=this.shadowRoot.getElementById(\"func\");this.ioElem=this.shadowRoot.getElementById(\"io\");this.levelElem=this.shadowRoot.getElementById(\"level\");this.ioElem.addEventListener(\"change\",e=>{log(this.id+\".io.change(\"+this.ioElem.checked+\")\");if(this.ioElem.checked){Controller.setGPIOOut(this.pin.gpio,this.levelElem.checked);}else{Controller.setGPIOIn(this.pin.gpio);}},{passive:true});this.levelElem.addEventListener(\"change\",e=>{log(this.id+\".level.change(\"+this.levelElem.checked+\")\");if(this.ioElem.checked){Controller.setGPIOOut(this.pin.gpio,this.levelElem.checked);}},{passive:true});}\nsetupPin(pin){this.pin=pin;this.id=this.pin.name;this.shadowRoot.getElementById(\"name\").textContent=this.pin.name;if(this.pin.gpio){log(this.id+\" is GPIO\");this._isGPIO();return;}\nlog(this.id+\" is indeterminate\");if(this.pin.func){this.funcElem.textContent=this.pin.func;this.funcElem.style.display=\"inline-block\";}\nController.eventGPIO.addEventListener(this.pin.name,()=>{log(this.id+\" is GPIO (late)\");this._isGPIO();},{once:true});}\n_isGPIO(){Controller.eventGPIO.addEventListener(this.pin.name,()=>this._gpioUpdate());this.shadowRoot.querySelector(\"div\").classList.add(\"gpio\");this._gpioUpdate();}\n_gpioUpdate(){log(this.id+\"._gpioUpdate(\"+this.pin.func+\")\")\nif(this.pin.func.startsWith(\"In/\")||this.pin.func.startsWith(\"Out/\")){this.funcElem.textContent=\"\";this.funcElem.style.display=\"none\";this.shadowRoot.querySelector(\".controls\").style.display=\"inline-block\";this.ioElem.checked=this.pin.func.startsWith(\"Out/\");this.ioElem.disabled=false;this.ioElem.indeterminate=false;this.levelElem.checked=this.pin.func.endsWith(\"/High\");this.levelElem.disabled=!this.ioElem.checked;if(this.ioElem.checked){this.ioElem.text=\"Out\";}else{this.ioElem.text=\"In\";}\nif(this.levelElem.checked){this.levelElem.text=\"High\";this.levelElem.indeterminate=false;}else if(this.pin.func.endsWith(\"/Low\")){this.levelElem.text=\"Low\";this.levelElem.indeterminate=false;}else{this.levelElem.indeterminate=true;this.levelElem.text=\"Ind\";}}else{if(this.pin.func){this.funcElem.textContent=this.pin.func;this.funcElem.style.display=\"inline-block\";}else{this.funcElem.textContent=\"\";this.funcElem.style.display=\"none\";}\nthis.ioElem.checked=false;this.ioElem.disabled=true;this.ioElem.indeterminate=true;this.ioElem.text=\"I/O\";this.levelElem.checked=false;this.levelElem.disabled=true;this.levelElem.indeterminate=true;this.levelElem.text=\"Level\";}}});</script><template id=template-header-view><data-table-elem></data-table-elem></template><script>\"use strict\";window.customElements.define(\"header-view\",class extends HTMLElementTemplate{constructor(){super(\"template-header-view\");}\nsetupHeader(name){this.header=Controller.headers[name];let data=this.shadowRoot.querySelector(\"data-table-elem\");let cols=1;if(this.header.pins){cols=this.header.pins[0].length;}\nlet hdr=[this.header.name];for(let i=1;i<cols;i++){hdr[i]=\"\";}\ndata.setupTable(hdr);for(let y=0;y<this.header.pins.length;y++){let row=this.header.pins[y];let items=[];for(let x=0;x<row.length;x++){items[x]=document.createElement(\"gpio-view\");}\nitems=data.appendRow(items);for(let x=0;x<items.length;x++){items[x].setupPin(row[x]);}}}});</script><template id=template-i2c-elem><data-table-elem></data-table-elem></template><script>\"use strict\";window.customElements.define(\"i2c-elem\",class extends HTMLElementTemplate{constructor(){super(\"template-i2c-elem\");}\nsetupI2C(name,number,err,scl,sda){let data=this.shadowRoot.querySelector(\"data-table-elem\");data.setupTable([name,\"\"]);if(number!=-1){data.appendRow([\"Number\",number]);}\nif(err){data.appendRow([\"Error\",err]);}\nif(scl){data.appendRow([\"SCL\",scl]);}\nif(sda){data.appendRow([\"SDA\",sda]);}}});</script><template id=template-spi-elem><data-table-elem></data-table-elem></template><script>\"use strict\";window.customElements.define(\"spi-elem\",class extends HTMLElementTemplate{constructor(){super(\"template-spi-elem\");}\nsetupSPI(name,number,err,clk,mosi,miso,cs){let data=this.shadowRoot.querySelector(\"data-table-elem\");data.setupTable([name,\"\"]);if(number!=-1){data.appendRow([\"Number\",number]);}\nif(err){data.appendRow([\"Error\",err]);}\nif(clk){data.appendRow([\"CLK\",clk]);}\nif(mosi){data.appendRow([\"MOSI\",mosi]);}\nif(mosi){data.appendRow([\"MISO\",miso]);}\nif(cs){data.appendRow([\"CS\",cs]);}}});</script><div class=err id=err></div><h1>GPIO</h1><div id=section-gpio></div><div id=section-state><h1>periph's state</h1><div><drivers-elem id=section-drivers-loaded></drivers-elem><drivers-elem id=section-drivers-skipped></drivers-elem><drivers-elem id=section-drivers-failed></drivers-elem></div></div><h1>I²C</h1><div id=section-i2c></div><h1>SPI</h1><div id=section-spi></div>"),
	"static/favicon.ico": []byte("\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR\x00\x00\x00\x80\x00\x00\x00\x80\b\x06\x00\x00\x00\xc3>a\xcb\x00\x00\ahIDATx\xda\xed\x9d\xcdkTW\x18\xc6\u007f9\x99aH\x98\xd1$^\xb5\"\xa6%Q\xf0\x83\ba6ҍ\xbb\xd6t\xd3\xddh5b\x15\xbb(T\xd4R\x84\xe4\x0fH@\nV\\\xeaF\x1c\xad\xceZ\xd0v\xe7\xce\xcdm\xc0\xc1\x0f*\t4RDs\xd5\xc4\f\xd1a\x92I\x17g\xc0\xceG\xe6+\xc9\xcc\xdcs\xde\xdfr`\xe6\xdey\x9e\xe7\xbe\xe7\xdcsν\xa7\x8d:\x89\xc7\xe3˅\x9f\r\x0f\x0f\xb7\xd1 \xca\x1dߋ\xc6:\xa1\xad\x17\xd4\x00\xb4\x0fB\xfb^\b\xf6Ap+\x04\"\x10\fA@\x81\xca}3\v,f!\x93\x86\xc5yȼ\x82\xcc\x14,=\x81\xa5\t\xc8&ay\xdaq\x13\v~\xf8\xff\xb5\x10\xc0 \xbc\xe8\xe9G\xb0a\x1f\x84Tm\xdfl\a\xda\x15\x84:\x80\x0e`\v0\x00|\x9b\xff\xfb?g\xe1\xfdc\xf8x\xf7\xbe!\x9a\x05\xfcitlWi\x036\x0f\xac\xef\x91C*w\x8c\x01$\x00\x8d6\xfd\xf0n\b]\x80\x9e\xe3\xd0\x11h\xcds<\x97\x81\xb77 }\xd1q\xef<\x93\x00\xac\xfeJ\x0fC\xf0\ft\x8fB$\xdc\xfarv\x04`\xfbI\xe0\xa4\x17\xfd)\x05\xef\xc6 s\xc5q\x13)\t@\x8d%\x1e:/ö!\xff\x16\xd7H\x18\"c\xc0\x98\x17\xfd\xfe\x1e,\x9cu\xdc\xc4s\t@Y\xe3\x8f\fB\xf8\xfa\xfa\xb7\xe5\x8df\xdb\x100\xe4EO'!u\xc2qoOH\x00\x8a\xda\xf7H\xc2<\xe3KuR7\xff\xa5\x830\x1fk\x85~B\xa0\xb9\xc6Ǻ\xa0\xf3\x96\xbfK}\xddAx\x9ak\x1a\x8e:nb\xb6Yg\xa2\x9ag\xfe\xb1\x11\xe8\u007fg\x9f\xf9\x85MC\xff;\xad\x85%\x15@\x97\xfb\x9e\aе\x05!G\xef\x98\x17\xfd\xf1\x1c\xbc=\xd8\xe8fA5\xd6\xfc\xe11\xe8{*既k\v\xf4=\xd5\x1a\x19V\x01\xbch́\x8d\x0f\xc1\xe9\x17\xa3+\xb1cċ\xfe\x10\x83\xb9\x03\x8e\x9b\xf0|_\x01\xbc\xe8w_Aߌ\x98_\vN?\xf4\xcdh\xed֗\xb6R\xb3J\x82=(\x91@\x02 H\x00\x04\t\x80`%u/a*\xd5y<t\xe9q\xc3N\xfc\xfe\xf9}\xd8{\xfc\xb97\xf7\xcf\u007f\xb9\xa9\xf0\xd3z\x96\x84I\x05\xf0%\x1b7I\x13 H\x00\x04\t\x80Ќ\x004b\x88Rh\xd1\x00艝\xcf\xff\x10\xe9Z\x0f\xedͺW\x80\x8d\x0fWq\xf7(\xac\xef\xdd\xc1\xc3u\r\x80\x9e\xab\x96Y\xbd\xd6\xc5\xe9\xafu=\x81\xaa\xde\xfcûaǈ\x88\xdc\xea\xec\x18\xd1^\xady\x05\xe8y \xe2\xfa\x85\xea\xbdR\xd5]\xfd\xc7Fd\x19\x97\x9f\xe8\xdaR\xedBSU\xd9\xfcX\x17\U0010e268~\xa3wL{\xb7\xea\n\xd0yK\xc4\xf4+\x95\xbdS\x95;~6\xaf\xdb\xf7;ۆ*u\b+T\x80HBD\xf4;\xe5=T+_\xfdG\x06\xcd\u007fV\xcf\x066\x0fh/k\xae\x00\xe1\xeb\"\x9e)\xac\xec\xa5Z\xa1\xe7\xbfK\xae~Ӫ@lW\r\x15\xa0\xf3\xb2\x88f\xdc\x1d\xc1\xe5\xaa\x02\xa0_\xcb\"=\u007f3\xef\bb\xe1**@\xf0\x8c\x88e*\xc5ޖ\b@\xf7\xa8\be*\xc5ު\xfc\xf2\u007fx\xb7?\xde\xc6%\xd49&\x10.\x1c\x18*\xa8\x00\xa1\v\"\x92\xe9\xe4{\\\x10\x80\x9e\xe3\"\x90\xe9\xe4{\xac\xf2\xef\xfd;\x02\"\x90\xe9t\x04\xfe?&P\xf7\xfb\x01\x86\u007f9\xd6r\u007f-\xfe\xeb͆\x1d˔\xff/\xcf\x05X\x8e\x04@\x02 H\x00\x04k\xc9m\xb1r\xfaQ\xad\xb3\u007fο\xc5\xcb\x04\xbd\xed\x8d\x1bD\x94\xe3\xaf\xf6\xf83Iǽ\xb6?W\x016\xec\x93k\xc16\xb4\xe7Jo\xb0\x14\x92\xa6\xc0:Bʋ\xc6:\x95\xde]K\xb0\xb4\aЫ\xf4\xd6j\x82\xa5\xf7\x00\x03J\xef\xab'\xd8I\xfb\xa0қ*\n\x96\x06`\xaf\xd2;j\nv\x12\xecSz;U\xc1\xd2\x00lUz/]\xc1N\x02\x11\xa57R\x16,\xad\x00!\xa5w\xd1\x16,\xad\x00J\xc9|\x90\xd5\xe3\x00\xe2\xbeD\x80\xac\xa8`-Y\x14,J\x02\xace1\xab \x93\x16!l%\x93V\xb08/BX[\x01\xe6\x15d^\x89\x10\xd6V\x80W\n2S\"\x84\xb5\x01\x98R\xb0\xf4D\x84\xb0\x95\xa5'\n\x96&D\bk\x030\xa1 \x9b\x14!\xac\x1d\aH*X\x9e\x16!leyZ9nb\x01\xd22\x18d\x1d\xe9\xac\xe3&\x16rs\x01\xef\x1f\x8b \xb6\xa1=\xcf\x05\xe0\xe3]\x11\xc46\xb4\xe7m|\xb6\xbc\xbcV?)\xcf\xe7\xfb\xef\xff\xcbt\xb0\xe5H\x00$\x00\x82\x04@\xb0\x96\xbc\x1d \xbd\xe8\xb9L\xb5o\n\xbb\u007f\xbe\xf8\x89\xf2C\x97\x1aw7)ǯ\xf7\xf8\x1f\x16\x1d\xf7\xb7\xe0\n\x15\xe0\xed\r\xb9&L'\xdf\xe3\x82\x00\xa4/\x8a@\xa6\x93\xefq^\x00\x1c\xf7\xce3\x98O\x89H\xa62\x9f\xd2\x1e\x97\xed\x04\xbe\x93=\x02\x8d\xa5\xd8\xdb\x12\x01\xc8\\\x11\xa1L\xa5\xd8ۢ\x008n\"\x05/\xef\x89X\xa6\xf1\xf2\x9e\xf6\xb6\xaaq\x80\x85\xb3\"\x98i\x94\xf6\xb4d\x00\x1c7\xf1\x1cfd\xa5\x901\xcc$\xb5\xa7U\x06@\x93:!\u0099\xc2\xca^\xae\x18\x00ǽ=!U\xc0\x94\xab\xff\xf6D\xcd\x01\xc8\xdd7\xc6D@\xdf\xdf\xfb\x97\xf5\xb0l\x00\xf4\xa0\x81\xdc\x11\xf8\xbb\xe7\x9f?\xf0Sc\x05\x00X8*B\xfa\xb6\xe7_ѻ\x8a\x01p\xdc\xc4,L\xcb^\x82\xbeczT{\xb7\xca\x00\xe8\x10\xdc\x1c\x87\xd9\xd7\"\xaa_\x98}\xad=\xabL\r\vB\xde\x1e\x14a\xfdB\xf5^U\x1d\x00ݙx1.\xe2\xb6:/\xc6+u\xfc\xea\xac\x00\xe0\xb8\xf1Q\xf0&E\xe4Vś\xd4\x1eUO\x1dk\x02\xe7\x0e\xc0\xb2hݒ\xcc\x1d\xa8\xf5\x1b5\a\xc0q\x13\x1e\xfc\xf3\xb5\x88\xddzho\xd69\x00\xfa@\xbf\xff)r\x9b\x81,\v\x97\x00\b\x12\x00A\x02 \xf8\xaa\xb7\xff\xa6E\x03\xb0v'&\x94\xd3\xd8\xfbb\xad~\xad-\x1e\x8f\xcbM\xbd4\x01\x82\x04@\x90\x00\b\x12\x00\xc12\xda\xea\xfdb\xa9\xce\xe3\xf0\xf0p\xc9\xdf\xf3\xa21\a6>\x04\xa7\u007f\xadN\xdc\xec\xf7\x03x\x930w\xa0\xdc\xd8~-\xfa7\xbd\x028n\xc2sܫ;e=A5\xbc\x18wܫ;\xeb\x99\xd8i\xf9&@\xcfUO\xed\x91\xe5e\xa5\x98}\rS{j\x9d\xcf_-\x81F\xff\xcd\xdcj\x95\xad^\xf4\xd8\b\xf4ʣ\xe8\x80^\xc0y\xb3)ձi\x9d@\xfd\x87'\xbb\xed~\xee\xe0\xe5=\x98\xecn\x96\xf9M\xa9\x00\x05}\x83Y\xe0\x1b/zx7D\x12\xb0y\xc0\x0e\xe3g\x920\x1f\xabe힑\x01(h\x16\xf6{\xd1#\x83\x10\xbenn\x10f\x92\x90:Q\xeeY=+\x03\xf0)\b\xb7't\x10b\xbb\xa0\xf32l\x1b2\xa7\xd4/\x9c]\xe9\x11m\t@q\xd3\xf0\\7\r\xb10\x04\xcf@\xf7(D\xc2\xfe2}>\xa5\xdfɓ\xb9R\xea\xcd\x1c\x12\x80ꂐ\x02Ɓq\xddO\b]\x80\x9e\xe3վ̲\xf1|X\xd4\xef\xe1K_l\x85\xf6\xdd\xf7\x01(\xd1O8\x05\x9c\xd2Mľ\xbf[\xef\x1c?\xbd\x81\xd3/\xf8r.`\xe5\xb6t&\xb9\xbe\xdbߤ\xb3\xfa\x18\xe6\x8ch\x060\bǽ\xb6\x1f\xc0\x8b\xc6:\xa1\xad\x17\xd4\x00\xb4\x0fB\xfb^\b\xf6Ap+\x04\"\x10\fA@}\xca\u007f\x16\xbd\x89v&\xad\xb7\xd2ͼ\xd2\x1bj.=\xd1\xdb\xeae\x93\xb0<\xad\xf7W\xd2\xc494b\x82f\xff\x01\xf7Qi\xbd}\xf6\x1b\xc6\x00\x00\x00\x00IEND\xaeB`\x82"),
}
