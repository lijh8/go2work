package main

/*
#include <stdlib.h>
#include <string.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {

	// cgo: pass strings to and from C function
	// char *strncpy2(char *dest, const char *src, size_t n);

	buf := [32]byte{}
	buflen := len(buf)
	str := "hello"
	cs := C.CString(str)
	defer C.free(unsafe.Pointer(cs))

	C.strncpy(
		(*C.char)(unsafe.Pointer(&buf)),
		cs,
		C.size_t(buflen)-1)

	fmt.Println(string(buf[:]))

}
