package main

import (
    "net/http"
    "time"
    "github.com/gorilla/websocket"
	"log"
	"net"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool { return true },
}

func close(conn *websocket.Conn){
    conn.WriteControl(
        websocket.CloseMessage,
        websocket.FormatCloseMessage(
            websocket.CloseGoingAway,
            ""), 
        time.Now().Add(time.Second))
}

func send(conn *websocket.Conn, message []byte) (bool){
    err := conn.WriteMessage(websocket.TextMessage, message)
    if err != nil {
        return false
    }
    return true
}

func readLoop(c *websocket.Conn) {
    for {
		_, _, err := c.NextReader();
        if err != nil {
            c.Close()
            break
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
	l, err := net.ListenMulticastUDP("udp", nil, addr)
	l.SetReadBuffer(1500)
	for {
		b := make([]byte, 1500)
		_, _, err := l.ReadFromUDP(b)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}
		err = conn.WriteMessage(websocket.TextMessage, b)
		if err != nil {
			break
		}
	}

}