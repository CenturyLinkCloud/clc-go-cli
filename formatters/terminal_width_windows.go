package formatters

import (
	"syscall"
	"unsafe"
)

type (
	screenBufferInfo struct {
		size                coord
		cursor_position     coord
		attributes          uint16
		window              small_rect
		maximum_window_size coord
	}
	coord struct {
		x int16
		y int16
	}
	small_rect struct {
		left   int16
		top    int16
		right  int16
		bottom int16
	}
)

var DEFAULT_WIDTH = 80

var kernel32 = syscall.NewLazyDLL("kernel32.dll")

var (
	procGetConsoleScreenBufferInfo = kernel32.NewProc("GetConsoleScreenBufferInfo")
)

func SetTerminalWidthFn(f func() uint) {
	getTerminalWidth = f
}

func getConsoleScreenBufferInfo(h syscall.Handle) (info screenBufferInfo, err error) {
	info = screenBufferInfo{}
	r0, _, e1 := syscall.Syscall(procGetConsoleScreenBufferInfo.Addr(),
		2, uintptr(h), uintptr(unsafe.Pointer(&info)), 0)
	if int(r0) == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

var getTerminalWidth = func() uint {
	out, err := syscall.Open("CONOUT$", syscall.O_RDWR, 0)
	if err != nil {
		return uint(DEFAULT_WIDTH)
	}
	info, err := getConsoleScreenBufferInfo(out)
	if err != nil {
		return uint(DEFAULT_WIDTH)
	}
	return uint(info.size.x)
}
