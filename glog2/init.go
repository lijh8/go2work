package glog2

import (
	"flag"
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

var Info = glog.Info
var Warning = glog.Warning
var Error = glog.Error
var Fatal = glog.Fatal

func init() {
	// provide an existing directory as log_dir flag at command line,
	// $ ./hello -log_dir=/path/to/logs
	// it create a subdir with timezone if log_dir is not provided.
	timezone, _ := time.Now().Zone()
	dirname := "LOG"
	dirname += "_" + timezone
	exePath, _ := os.Executable()
	realPath, _ := filepath.EvalSymlinks(exePath)
	log_dir := filepath.Dir(realPath)
	log_dir = filepath.Join(log_dir, dirname)

	flag.Parse()
	if flag.Lookup("log_dir").Value.String() == "" {
		os.MkdirAll(log_dir, 0755)
		flag.Set("log_dir", log_dir)
	}
	glog.MaxSize = 1024 * 1024 * 5 // 1.8G default
}

/*
import (
	. "glog2" // https://golang.google.cn/ref/spec#Import_declarations

	"github.com/golang/glog"
)

func main() {
	defer glog.Flush()
	Info()
	Warning()
	Error()
}
*/

/*
$ go build . && ./main
$ go build . && ./main -log_dir=LOG_CST

$ cat ./LOG_CST/glog.ljhs-Mac-mini.ljh.log.INFO.20240930-003828.19717
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
