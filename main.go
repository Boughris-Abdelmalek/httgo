package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// net.Listen accept a connection type (tcp in this case) with a port an address
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// go starts an async process
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	fmt.Printf("Serving %s\n", conn.RemoteAddr().String())
	packet := make([]byte, 1024)
	// we use a buffer to temporarily store the data until we can process it
	buffer := make([]byte, 1024)
	// we use defer to close the connection in the loop to handle the incoming data
	defer conn.Close()

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error: ", err)
			}
			break
		}
		// we append the buffer's bytes to a single packet
		packet = append(packet, buffer...)

		fmt.Printf("Received: %s\n", buffer[:n])
	}
	num, _ := conn.Write(packet)
	fmt.Printf("Wrote back %d bytes, the payload is %s\n", num, string(packet))
}
