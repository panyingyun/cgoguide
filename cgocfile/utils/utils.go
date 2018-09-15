package utils

/*
#cgo CFLAGS: -std=c99 -I ./
#cgo LDFLAGS: -L.
#include "utils.h"
*/
import "C"

import "fmt"

func GoSum(a, b int) int {
	s := C.sum(C.int(a), C.int(b))
	fmt.Println("GoSum s =", s)
	return int(s)
}
