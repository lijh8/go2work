package main

import (
	"bufio"
	"fmt"
	"log"
	"log2"
	"net"
	"os"
)

// $ cd client
// $ go build && ./client 192.168.1.3 12345 clientABC

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
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	done := make(chan any)

	go func() {
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				log.Println(err, line)
				done <- nil
				break
			}
			fmt.Print(line)
		}
	}()

	go func() {
		for {
			msgID++
			msg := fmt.Sprintf("msg from client: %s, msgID: %d\n", tag, msgID)
			if n, err := writer.WriteString(msg); err != nil {
				log.Println(err, n)
				done <- nil
				break
			}
			writer.Flush()
		}
	}()

	<-done
	<-done
}
