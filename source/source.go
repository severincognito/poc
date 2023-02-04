package main

import (
	"encoding/xml"
	"log"
	"net"
	"time"
)

type Message struct {
	Title     string `xml:"title"`
	Timestamp int64  `xml:"timestamp,attr"`
	Content   string `xml:"content"`
}

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:5000")
	if err != nil {
		log.Fatal(err)
		return
	}

	for {
		// marshal and send it to udp server..
		msg, err := xml.Marshal(Message{Title: "bonjour de xml", Content: "123", Timestamp: time.Now().UnixMilli()})
		if err != nil {
			log.Fatal(err)
			return
		}
		conn.Write(msg)
		time.Sleep(1000 * time.Millisecond)
	}

	conn.Close()
}
