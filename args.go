package main

import (
    "github.com/docopt/docopt-go"
)

var inifile string

func args_init() {
    usage := `Websocket API proxy

Usage:
  ws_proxy [--cfg=<inifile>]

Options:
  --cfg=<inifile>    Settings file path [default: settings.ini]

`
    arguments, _ := docopt.ParseDoc(usage)
    inifile = arguments["--cfg"].(string)
}