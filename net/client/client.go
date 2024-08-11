package main

import (
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

	if len(os.Args) != 4 {
		log.Printf("Usage: ./client <serverIP> <port> <tag>")
		os.Exit(-1)
	}

	serverIP := os.Args[1]
	port := os.Args[2]
	tag := os.Args[3]
	conn, err := net.Dial("tcp", serverIP+":"+port)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	handleConnection(conn, tag)
}

func handleConnection(conn net.Conn, tag string) {
	defer conn.Close()

	msgID := 0
	buffer := make([]byte, 1024)

	for {
		msgID++
		msg := fmt.Sprintf("msg from client: %s, msgID: %d", tag, msgID)
		n, err := conn.Write([]byte(msg))
		if err != nil {
			log.Println(err, n)
			break
		}

		n, err = conn.Read(buffer)
		if err != nil {
			log.Println(err, n)
			break
		}
		fmt.Println(string(buffer))
	}
}
