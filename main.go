package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	_ "github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcapgo"
	"log"
	"net"
	"os"
)

type AppMessage struct {
	Counter int64
	Message Message
}

type Message struct {
	Title     string `xml:"title"`
	Timestamp int64  `xml:"timestamp,attr"`
	Content   string `xml:"content"`
}

func savePcap() {
	f, err := os.Create("lo.pcap")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	pcapw := pcapgo.NewWriter(f)
	if err := pcapw.WriteFileHeader(1600, layers.LinkTypeEthernet); err != nil {
		log.Fatalf("WriteFileHeader: %v", err)
	}

	handle, err := pcapgo.NewEthernetHandle("lo")
	if err != nil {
		log.Fatalf("OpenEthernet: %v", err)
	}

	pkgsrc := gopacket.NewPacketSource(handle, layers.LayerTypeEthernet)
	for packet := range pkgsrc.Packets() {
		if err := pcapw.WritePacket(packet.Metadata().CaptureInfo, packet.Data()); err != nil {
			log.Fatalf("pcap.WritePacket(): %v", err)
		}
	}
}

func getXML(c chan Message) {
	addr := net.UDPAddr{
		Port: 5000,
		IP:   net.ParseIP("127.0.0.1"),
	}

	s, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Println(err)
		return
	}

	b := make([]byte, 512)
	for {
		// 1. read it to byte slice
		n, _, err := s.ReadFrom(b)

		//fmt.Println(string(b))

		// 2. unmarshal it to Message struct
		var msg Message
		if err = xml.Unmarshal(b[:n], &msg); err != nil {
			log.Println(err)
			continue
		}
		c <- msg
	}
}

func main() {
	var counter int64

	//go savePcap()

	input := make(chan Message, 10)
	go getXML(input)

	app := fiber.New()
	app.Use(cors.New())
	//app.Static("/", "./static")

	app.Use("/resetCounter", func(c *fiber.Ctx) error {
		counter = 0
		return c.JSON("ok")
	})

	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
		for {
			m := <-input
			appM := AppMessage{counter, m}
			b, err := json.Marshal(appM)
			if err != nil {
				panic(err)
			}
			//fmt.Println(string(b))
			if err = c.WriteMessage(websocket.TextMessage, b); err != nil {
				log.Println("write:", err)
				break
			}
			counter++
		}
	}))

	app.Post("/transmit", func(c *fiber.Ctx) error {
		// getting json payload
		var payload Message
		if err := json.Unmarshal(c.Body(), &payload); err != nil {
			fmt.Println("oh no")
			return err
		}
		// converting to xml
		data, err := xml.Marshal(payload)
		if err != nil {
			fmt.Println("oh no")
			return err
		}
		// sending xml out...
		fmt.Println(string(data))

		return c.JSON("ok")
	})

	log.Fatal(app.Listen(":8080"))
}
