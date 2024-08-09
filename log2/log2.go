package log2

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func InitLog(filename string) (*os.File, error) {
	exePath, err := os.Executable()
	if err != nil {
		return nil, err
	}

	realPath, err := filepath.EvalSymlinks(exePath)
	if err != nil {
		return nil, err
	}

	logpath := filepath.Dir(realPath)
	logpath = filepath.Join(logpath, filename)
	logfd, err := os.OpenFile(logpath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	multiWriter := io.MultiWriter(logfd, os.Stdout)
	log.SetOutput(multiWriter)
	log.SetPrefix("UTC ")
	log.SetFlags(log.LUTC | log.LstdFlags | log.Lshortfile)
	return logfd, nil
}

/*
func main() {
	filename := "logfile.log"
	logfd, err := log2.InitLog(filename)
	if err == nil {
		defer logfd.Close()
	}

	log.Println("aaa")
	log.Println("bbb")
}
*/
