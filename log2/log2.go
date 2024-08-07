package log2

import (
	"io"
	"log"
	"os"
)

func InitLog(logfile string) *os.File {
	logfd, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	multiWriter := io.MultiWriter(logfd, os.Stdout)
	log.SetOutput(multiWriter)
	log.SetPrefix("UTC ")
	log.SetFlags(log.LUTC | log.LstdFlags | log.Lshortfile)
	return logfd
}
