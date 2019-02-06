package main

import (
	"fmt"
	"os"
	"os/signal"
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

	out, err = os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()

	signalCh := make(chan os.Signal)
	signal.Notify(signalCh, syscall.SIGWINCH)

	for {
		termw, termh := getTermSize(out.Fd())
		if termw > 0 && termh > 0 {
			fmt.Printf("termw: %d, termh: %d\n", termw, termh)
			return
		}

		select {
		case ch := <-signalCh:
			if ch == syscall.SIGWINCH {
				termw, termh := getTermSize(out.Fd())
				fmt.Printf("termw: %d, termh: %d\n", termw, termh)
				return
			}
		}
	}
}
