# AsciiPath
Simple golang web service that takes a string as input and returns ASCII art rendering (via https://github.com/common-nighthawk/go-figure). Using this as a personal hello world learning project for using WASM.

#### Usage

The Golang server will listen on whatever port is passed via `PORT` envvar; otherwise, it will default to `8080`.

Currently, the server responds to either:
- `GET`s will base the response on the path passed to the server. `/`s will be replaced with spaces (` `)
- `PUT`s will base the response on the string passed to the server 

```default
$ curl 127.0.0.1:8080/how/are/you
    _
   | |__     ___   __      __     __ _   _ __    ___     _   _    ___    _   _
   | '_ \   / _ \  \ \ /\ / /    / _` | | '__|  / _ \   | | | |  / _ \  | | | |
   | | | | | (_) |  \ V  V /    | (_| | | |    |  __/   | |_| | | (_) | | |_| |
   |_| |_|  \___/    \_/\_/      \__,_| |_|     \___|    \__, |  \___/   \__,_|
```

Set the font via the `font` header - available fonts listed [here](https://github.com/common-nighthawk/go-figure#supported-fonts)
```Testing a font via the font header
$ curl 127.0.0.1:8080/how/are/you --header "font: cosmic"
    ::   .:      ...     .::    .   .:::    :::.     :::::::..   .,::::::    .-:.     ::-.    ...      ...    :::
   ,;;   ;;,  .;;;;;;;.  ';;,  ;;  ;;;'     ;;`;;    ;;;;``;;;;  ;;;;''''     ';;.   ;;;;' .;;;;;;;.   ;;     ;;;
  ,[[[,,,[[[ ,[[     \[[, '[[, [[, [['     ,[[ '[[,   [[[,/[[['   [[cccc        '[[,[[['  ,[[     \[[,[['     [[[
  "$$$"""$$$ $$$,     $$$   Y$c$$$c$P     c$$$cc$$$c  $$$$$$c     $$""""          c$$"    $$$,     $$$$$      $$$
   888   "88o"888,_ _,88P    "88"888       888   888, 888b "88bo, 888oo,__      ,8P"`     "888,_ _,88P88    .d888
   MMM    YMM  "YMMMMMP"      "M "M"       YMM   ""`  MMMM   "W"  """"YUMMM    mM"          "YMMMMMP"  "YmmMMMM""
```

A `PUT` example
```
$ curl -X PUT http://127.0.0.1:8080 -d "123   45" --Header "font: cosmic"
:.  .:::.   .::.             .:: ::::::::
;; ,;'``;. ;'`';;,         ,;';; `;;``'';
[[ ''  ,[['   .n[[        ,[' [[  [[,_
$$ .c$$P'    ``"$$$.      $P__$$c `""*Ycc
88d88 _,oo,  ,,o888"      `"""88" __,od8"
MMMMMUP*"^^  YMMP"            MM  MMP"
```

Examples when using Cloud Run
```
$ curl https://asciipath-4uotx33u2a-uc.a.run.app/Hello/World --Header "font:banner3-D"
'##::::'##:'########:'##:::::::'##::::::::'#######::  '##:::::'##::'#######::'########::'##:::::::'########::
 ##:::: ##: ##.....:: ##::::::: ##:::::::'##.... ##:   ##:'##: ##:'##.... ##: ##.... ##: ##::::::: ##.... ##:
 ##:::: ##: ##::::::: ##::::::: ##::::::: ##:::: ##:   ##: ##: ##: ##:::: ##: ##:::: ##: ##::::::: ##:::: ##:
 #########: ######::: ##::::::: ##::::::: ##:::: ##:   ##: ##: ##: ##:::: ##: ########:: ##::::::: ##:::: ##:
 ##.... ##: ##...:::: ##::::::: ##::::::: ##:::: ##:   ##: ##: ##: ##:::: ##: ##.. ##::: ##::::::: ##:::: ##:
 ##:::: ##: ##::::::: ##::::::: ##::::::: ##:::: ##:   ##: ##: ##: ##:::: ##: ##::. ##:: ##::::::: ##:::: ##:
 ##:::: ##: ########: ########: ########:. #######::  . ###. ###::. #######:: ##:::. ##: ########: ########::
..:::::..::........::........::........:::.......:::  :...::...::::.......:::..:::::..::........::........:::
```

A `PUT` example to Cloud Run with an empty payload
```
$ curl -X PUT https://asciipath-4uotx33u2a-uc.a.run.app -d "" --Header "font: cosmic"
.,::::::  .        :   ::::::::::. ::::::::::::.-:.     ::-.  :::::::.      ...     :::::::-.  .-:.     ::-.
;;;;''''  ;;,.    ;;;   `;;;```.;;;;;;;;;;;'''' ';;.   ;;;;'   ;;;'';;'  .;;;;;;;.   ;;,   `';, ';;.   ;;;;'
 [[cccc   [[[[, ,[[[[,   `]]nnn]]'      [[        '[[,[[['     [[[__[[\.,[[     \[[, `[[     [[   '[[,[[['
 $$""""   $$$$$$$$"$$$    $$$""         $$          c$$"       $$""""Y$$$$$,     $$$  $$,    $$     c$$"
 888oo,__ 888 Y88" 888o   888o          88,       ,8P"`       _88o,,od8P"888,_ _,88P  888_,o8P'   ,8P"`
 """"YUMMMMMM  M'  "MMM   YMMMb         MMM      mM"          ""YUMMMP"   "YMMMMMP"   MMMMP"`    mM"
```