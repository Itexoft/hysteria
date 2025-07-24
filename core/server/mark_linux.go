//go:build linux

package server

import (
    "net"
    "golang.org/x/sys/unix"
    "syscall"
)

func setSocketMark(c net.Conn, mark uint32) {
    if mark == 0 {
        return
    }
    if sc, ok := c.(interface{ SyscallConn() (syscall.RawConn, error) }); ok {
        rc, _ := sc.SyscallConn()
        rc.Control(func(fd uintptr) {
            unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_MARK, int(mark))
        })
    }
}
