package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/golang/glog"
)

// $ cd /path/to/mymodule
// $ go get github.com/golang/glog  # it will modify go.mod, go.sum
// $ go mod tidy  # it will modify go.sum

// $ go build
// $ go run .
// $ go run . -log_dir=/path/to/logs
// $ ./hello
// $ ./hello -log_dir=/path/to/logs
// $ go clean

func InitLog() {
	// if user provides log_dir flag at command line,
	// $ ./hello -log_dir=/path/to/logs
	// the directory should exist.
	// create a subdir with timezone name if user does not provide one.
	timezone, _ := time.Now().Zone()
	dirname := "logs"
	dirname += "_" + timezone
	exePath, _ := os.Executable()
	realPath, _ := filepath.EvalSymlinks(exePath)
	log_dir := filepath.Dir(realPath)
	log_dir = filepath.Join(log_dir, dirname)

	flag.Parse()
	if flag.Lookup("log_dir").Value.String() == "" {
		if err := os.MkdirAll(log_dir, 0755); err != nil {
			if !os.IsExist(err) {
				fmt.Println(err)
				return
			}
		}
		flag.Set("log_dir", log_dir)
	}
	glog.MaxSize = 1024 * 1024 * 5 // 1.8 G default, for test
}

/*
$ cat ./logs_CST/glog.ljhs-Mac-mini.ljh.log.INFO.20240930-003828.19717
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
