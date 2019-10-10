package main

/*
#cgo CFLAGS: -I../include
#cgo LDFLAGS: -L../lib -lvideo
#include "video.h"
*/
import "C"

func main() {
	cmd := C.CString("ffmpeg -i ./xxx/*.png ./xxx/yyy.mp4")
	C.exeFFmpegCmd(&cmd)
}
