package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/golang/glog"
)

// $ cd /path/to/mymodule
// $ go get github.com/golang/glog
// $ go mod tidy # generates go.sum

func main() {
	// log to subdir with the executable
	subDir := "./logs"
	exePath, _ := os.Executable()
	realPath, _ := filepath.EvalSymlinks(exePath)
	parentDir := filepath.Dir(realPath)
	logDir := filepath.Join(parentDir, subDir)

	if err := os.MkdirAll(logDir, 0755); err != nil {
		if !os.IsExist(err) {
			glog.Error(err)
			return
		}
	}

	// set flags
	flag.Set("log_dir", logDir)
	flag.Set("stderrthreshold", "ERROR")

	flag.Parse()
	glog.MaxSize = 5 * 1024 * 1024 // in bytes
	defer glog.Flush()

	for {
		glog.Info("INFO!")
		glog.Warning("WARNING!")
		glog.Error("ERROR!")
		// glog.Fatal("FATAL!") // this will exit the program
	}
}
