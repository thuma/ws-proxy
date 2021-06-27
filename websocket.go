package main

import (
	"log"
	"net"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func send(conn *websocket.Conn, message []byte) bool {
	err := conn.WriteMessage(websocket.TextMessage, message)
	return err == nil
}

func readLoop(c *websocket.Conn) {
	for {
		_, _, err := c.NextReader()
		if err != nil {
			c.Close()
			return
		}
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	go readLoop(conn)
	addr, err := net.ResolveUDPAddr("udp", "224.0.0.1:9999")
	if err != nil {
		log.Fatal(err)
	}
	iface, err := net.InterfaceByName(ifacename)
	if err != nil {
		log.Println(err)
		return
	}
	l, _ := net.ListenMulticastUDP("udp", iface, addr)
	l.SetReadBuffer(1500)
	for {
		messagedata := make([]byte, 1500)
		_, _, err := l.ReadFromUDP(messagedata)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}
		if send(conn, messagedata) {
			conn.Close()
			return
		}
	}

}
