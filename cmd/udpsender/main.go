package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":42069")
	if err != nil {
		log.Fatal("resolve error:", err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("dial error:", err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println("read error:", err)
			continue
		}

		_, err = conn.Write([]byte(line))
		if err != nil {
			log.Println("write error:", err)
			continue
		}
	}
}
