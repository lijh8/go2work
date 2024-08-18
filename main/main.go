package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	n, err := writer.WriteString("Enter a line: ")
	if err != nil {
		fmt.Println(err, n)
		return
	}
	writer.Flush()

	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err, line)
		return
	}

	msg := fmt.Sprintf("You entered: %s\n", line)
	n, err = writer.WriteString(msg)
	if err != nil {
		fmt.Println(err, n)
		return
	}
	writer.Flush()
}
