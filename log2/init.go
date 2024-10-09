package log2

// multiple source files in same package;
// names from different source files in same package can be used without export;

import "log"

var LOG = log.Println
var LOGF = log.Printf

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
