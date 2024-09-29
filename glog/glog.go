package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang/glog"
)

// $ cd /path/to/mymodule
// $ go get github.com/golang/glog  # it will modify go.mod, go.sum
// $ go mod tidy  # it will modify go.sum

// $ go build
// $ ./hello -log_dir=/path/to/logs
// $ go clean

func main() {
	// if user provides log_dir, it should exist.
	// if user does not, a subdir is created.
	subDir := "./logs"
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
	glog.MaxSize = 1024 * 1024 * 5 // 1.8 G default, for test

	defer glog.Flush()

	for {
		glog.Info("INFO!")
		glog.Warning("WARNING!")
		glog.Error("ERROR!")
		// glog.Fatal("FATAL!") // this will exit the program
	}
}

/*
$ cat glog.ljhs-Mac-mini.ljh.log.INFO.20240930-003828.19717
Log file created at: 2024/09/30 00:38:28
Running on machine: ljhs-Mac-mini
Binary: Built with gc go1.23.1 for darwin/amd64
Previous log: logs/glog.ljhs-Mac-mini.ljh.log.INFO.20240930-003826.19717
Log line format: [IWEF]mmdd hh:mm:ss.uuuuuu threadid file:line] msg
W0930 00:38:28.781964   19717 glog.go:44] WARNING!
E0930 00:38:28.795429   19717 glog.go:45] ERROR!
I0930 00:38:28.795465   19717 glog.go:43] INFO!
W0930 00:38:28.795471   19717 glog.go:44] WARNING!
$
*/
