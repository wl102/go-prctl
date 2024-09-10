package prctl

import (
	"syscall"
	"unsafe"
)

func GetProcessName() (prname string, err error) {
	// linux limit max 16 bytes
	buf := make([]byte, 16)
	_, _, errno := syscall.Syscall(uintptr(syscall.SYS_PRCTL), uintptr(syscall.PR_GET_NAME), uintptr(unsafe.Pointer(&buf[0])), 0)
	if errno != 0 {
		err = errno
	} else {
		prname = string(buf[:15])
	}
	return
}

func SetProcessName(newname string) error {
	// c formate string
	buf := make([]byte, len(newname)+1)
	_ = copy(buf, newname)
	namePtr := uintptr(unsafe.Pointer(&buf[0]))
	_, _, errno := syscall.Syscall(uintptr(syscall.SYS_PRCTL), uintptr(syscall.PR_SET_NAME), namePtr, 0)
	if errno != 0 {
		return errno
	}
	return nil
}
