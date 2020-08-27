package main

import "fmt"

/*
#cgo CFLAGS: -I${SRCDIR}/../system
#include <cstuff/cdata.h>
*/
import "C"

func main() {
	text := C.GoString(C.the_text)
	fmt.Println(text)
}
