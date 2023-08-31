package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	// 1. 2. create and bind
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	// 3. listen
	ln, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	// log.Println("normal socket\nlisten on port", service)
	log.Println("concurrency socket\nlisten on port", service)
	for {
		// 4. accept
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleClient(conn)
		// read from socket
		req := make([]byte, 1024)
		len, err := conn.Read(req)
		log.Println("request: ", string(req[:len]))

		// // write to socket
		// daytime := time.Now().String()
		// conn.Write([]byte(daytime))
		// fmt.Println("conn write")

		// // 5. close
		// conn.Close()

	}

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	daytime := time.Now().String()
	conn.Write([]byte(daytime))
}
