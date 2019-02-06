package main

import (
	"fmt"
	"os"
	"runtime"
	"syscall"
	"unsafe"
)

type winsize struct {
	rows    uint16
	cols    uint16
	xpixels uint16
	ypixels uint16
}

func getTermSize(fd uintptr) (int, int) {
	var sz winsize
	_, _, _ = syscall.Syscall(syscall.SYS_IOCTL,
		fd, uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&sz)))
	return int(sz.cols), int(sz.rows)
}

func main() {
	var (
		err error
		out *os.File
	)

	if runtime.GOOS == "openbsd" {
		out, err = os.OpenFile("/dev/tty", os.O_RDWR, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		out, err = os.OpenFile("/dev/tty", os.O_WRONLY, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	termw, termh := getTermSize(out.Fd())

	fmt.Printf("termw: %d, termh: %d\n", termw, termh)
}
