# ascii-art-service
Simple golang web service that takes a string as input and returns ASCII art rendering (via https://github.com/common-nighthawk/go-figure). Using this as a personal hello world learning project for using WASM.

#### Usage

```default
$ curl 127.0.0.1:8080/how/are/you
    _
   | |__     ___   __      __     __ _   _ __    ___     _   _    ___    _   _
   | '_ \   / _ \  \ \ /\ / /    / _` | | '__|  / _ \   | | | |  / _ \  | | | |
   | | | | | (_) |  \ V  V /    | (_| | | |    |  __/   | |_| | | (_) | | |_| |
   |_| |_|  \___/    \_/\_/      \__,_| |_|     \___|    \__, |  \___/   \__,_|
```

```Testing a font via the font header
$ curl 127.0.0.1:8080/how/are/you --header "font: cosmic"
    ::   .:      ...     .::    .   .:::    :::.     :::::::..   .,::::::    .-:.     ::-.    ...      ...    :::
   ,;;   ;;,  .;;;;;;;.  ';;,  ;;  ;;;'     ;;`;;    ;;;;``;;;;  ;;;;''''     ';;.   ;;;;' .;;;;;;;.   ;;     ;;;
  ,[[[,,,[[[ ,[[     \[[, '[[, [[, [['     ,[[ '[[,   [[[,/[[['   [[cccc        '[[,[[['  ,[[     \[[,[['     [[[
  "$$$"""$$$ $$$,     $$$   Y$c$$$c$P     c$$$cc$$$c  $$$$$$c     $$""""          c$$"    $$$,     $$$$$      $$$
   888   "88o"888,_ _,88P    "88"888       888   888, 888b "88bo, 888oo,__      ,8P"`     "888,_ _,88P88    .d888
   MMM    YMM  "YMMMMMP"      "M "M"       YMM   ""`  MMMM   "W"  """"YUMMM    mM"          "YMMMMMP"  "YmmMMMM""
```

