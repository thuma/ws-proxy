package main

import (
	"log"

	"gopkg.in/ini.v1"
)

var http_server_port, ws_url, ifacename string

func read_cfg() {
	cfg, err := ini.Load(inifile)
	if err != nil {
		log.Fatal("Unable to read ini file.")
	}
	http_server_port = cfg.Section("").Key("http_server_port").String()
	ws_url = cfg.Section("").Key("ws_url").String()
	ifacename = cfg.Section("").Key("interface").String()
}
