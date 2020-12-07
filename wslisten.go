package main

import (
	"net"
	"github.com/gorilla/websocket"
	"log"
        "strings"
)

func wsListen() {
	addr, err := net.ResolveUDPAddr("udp", "224.0.0.1:9999")
	if err != nil {
		log.Fatal(err)
	}

        laddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9998")
        if err != nil {
                log.Fatal(err)
        }

	udpConn, _ := net.DialUDP("udp", laddr, addr)
	log.Println("Listen:", ws_url)
	wsconn, _, err := websocket.DefaultDialer.Dial(ws_url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer wsconn.Close()

	for {
		_, message, err := wsconn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		udpConn.Write([]byte(strings.Trim(string(message), "\n\r")))
	}
}
