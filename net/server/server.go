package main

import (
	"bufio"
	"fmt"
	"log"
	"log2"
	"net"
	"os"
	"time"
)

// $ cd server
// $ go build && ./server 12345

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

	const maxConcurrentConnections = 100000
	sem := make(chan struct{}, maxConcurrentConnections)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		clientAddr := conn.RemoteAddr().String()
		log.Println("accepted connection from", clientAddr)

		// if there is no read or write activity on the connection
		// for the specified duration, the connection will be closed.
		if err := conn.SetDeadline(time.Now().Add(3 * time.Minute)); err != nil {
			log.Println(err)
			conn.Close()
			continue
		}

		sem <- struct{}{}
		go func(c net.Conn) {
			defer func() { <-sem }()

			// the outer forever for loop in main keeps the
			// main goroutine alive and accepting new connections,
			// along with this closure goroutine.
			// no extra synchronization channel is needed here.
			handleConnection(c)
		}(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	msgID := 0
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	done := make(chan struct{})

	go func() {
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				log.Println(err, line)
				break
			}
			fmt.Print(line)
		}
		done <- struct{}{}
	}()

	go func() {
		for {
			msgID++
			msg := fmt.Sprintf("msg from server with msgID: %d\n", msgID)
			if n, err := writer.WriteString(msg); err != nil {
				log.Println(err, n)
				break
			}
			writer.Flush()
		}
		done <- struct{}{}
	}()

	_, _ = <-done, <-done
}
