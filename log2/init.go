package log2

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

// simple log with filename, line number source location
// 2024/10/09 20:41:59 main.go:11: abc

var LOG = log.Println
var LOGF = log.Printf

var logfd *os.File

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	logfile := "app.log"
	fd, err := InitLog(logfile)
	if err != nil {
		LOG(err)
		return
	}
	logfd = fd
}

func Close() {
	logfd.Close()
}

func InitLog(filename string) (*os.File, error) {
	exePath, _ := os.Executable()
	realPath, _ := filepath.EvalSymlinks(exePath)
	logpath := filepath.Dir(realPath)
	logpath = filepath.Join(logpath, filename)
	logfd, err := os.OpenFile(logpath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	multiWriter := io.MultiWriter(logfd, os.Stdout)
	log.SetOutput(multiWriter)
	return logfd, nil
}

/*

import (
	"log2"
	. "log2"
)

func main() {
	defer log2.Close()
	LOG()
	LOG("")
}

*/
