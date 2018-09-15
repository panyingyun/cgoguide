package main

/*

#cgo CFLAGS: -I../include

#cgo LDFLAGS: -L../lib -llibvideo

#include "video.h"

*/
import "C"

import "fmt"

func main() {
	fmt.Println("111")
	cmd := C.CString("ffmpeg -i ./xxx/*.png ./xxx/yyy.mp4")
	C.exeFFmpegCmd(cmd)
	fmt.Println("222")
}
