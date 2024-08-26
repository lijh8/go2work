package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang/glog"
)

// $ cd /path/to/mymodule
// $ go get github.com/golang/glog
// $ go mod tidy # generates go.sum

// $ go build
// $ ./hello -log_dir=/path/to/log

func main() {
	// if user provides log_dir, it should exist.
	// if user does not, a subdir is created.
	subDir := "./log"
	exePath, _ := os.Executable()
	realPath, _ := filepath.EvalSymlinks(exePath)
	logDir := filepath.Dir(realPath)
	logDir = filepath.Join(logDir, subDir)

	flag.Parse()
	if flag.Lookup("log_dir").Value.String() == "" {
		if err := os.MkdirAll(logDir, 0755); err != nil {
			if !os.IsExist(err) {
				fmt.Println(err)
				return
			}
		}
		flag.Set("log_dir", logDir)
	}
	glog.MaxSize = 1024 * 1024 * 5 // for test

	defer glog.Flush()

	for {
		glog.Info("INFO!")
		glog.Warning("WARNING!")
		glog.Error("ERROR!")
		// glog.Fatal("FATAL!") // this will exit the program
	}
}
