package main

import (
    "log"
    "net/http"
)

func main() {
    args_init()
    read_cfg()
    http.HandleFunc("/ws", wsHandler)
    go wsListen()
    log.Fatal(http.ListenAndServe(http_server_port, nil))
}