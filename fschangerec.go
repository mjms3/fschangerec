package main

import (
	"github.com/mjms3/fschangerec/errorhandling"
	"strings"
	"syscall"
	"unsafe"
)

func main() {

}

func monitor(pathname string, messageQueue chan []string) {
	fd, err := syscall.InotifyInit()
	errorhandling.FatalError(err)
	defer syscall.Close(fd)

	watchdesc, err := syscall.InotifyAddWatch(fd, pathname, syscall.IN_ALL_EVENTS)
	defer syscall.InotifyRmWatch(fd, uint32(watchdesc))
	errorhandling.FatalError(err)

	buffer := make([]byte, syscall.SizeofInotifyEvent*100)

	_, err = syscall.Read(fd, buffer)
	errorhandling.FatalError(err)
	raw := (*syscall.InotifyEvent)(unsafe.Pointer(&buffer[0]))
	bytes := (*[syscall.PathMax]byte)(unsafe.Pointer(&buffer[syscall.SizeofInotifyEvent]))
	name := strings.TrimRight(string(bytes[0:raw.Len]), "\000")
	messageQueue <- []string{name}

}
