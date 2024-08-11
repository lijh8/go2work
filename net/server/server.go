package main

import (
	"bufio"
	"fmt"
	"log"
	"log2"
	"net"
	"os"
)

func main() {
	logfile := "logfile.log"
	logfd, err := log2.InitLog(logfile)
	if err == nil {
		defer logfd.Close()
	}

	if len(os.Args) != 2 {
		log.Printf("Usage: ./server <port>")
		os.Exit(-1)
	}

	port := os.Args[1]
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			os.Exit(-1)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	msgID := 0
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Print(line)

		msgID++
		msg := fmt.Sprintf("msg from server with msgID: %d\n", msgID)
		if n, err := writer.WriteString(msg); err != nil {
			log.Println(err, n)
			break
		}
		writer.Flush()

	}
}
