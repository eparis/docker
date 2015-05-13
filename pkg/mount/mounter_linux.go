package mount

import (
	"fmt"
	"os"
	"syscall"
)

func mount(device, target, mType string, flag uintptr, data string) error {
	fmt.Fprintf(os.Stderr, "pkg/mount/mounter_linux.go syscall.Mount s=%s d=%s f=%x\n", device, target, flag)
	if err := syscall.Mount(device, target, mType, flag, data); err != nil {
		return err
	}

	// If we have a bind mount or remount, remount...
	if flag&syscall.MS_BIND == syscall.MS_BIND && flag&syscall.MS_RDONLY == syscall.MS_RDONLY {
		fmt.Fprintf(os.Stderr, "pkg/mount/mounter_linux.go syscall.Mount s=%s d=%s f=%x\n", device, target, flag|syscall.MS_REMOUNT)
		return syscall.Mount(device, target, mType, flag|syscall.MS_REMOUNT, data)
	}
	return nil
}

func unmount(target string, flag int) error {
	return syscall.Unmount(target, flag)
}
