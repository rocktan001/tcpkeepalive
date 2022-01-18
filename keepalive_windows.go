package tcpkeepalive

import (
	// "os"
	// "syscall"
	"time"
	// "unsafe"
)

func setIdle(c *Conn, secs int) error {
	c.TCPConn.SetKeepAlivePeriod(time.Duration(secs) * time.Second)
	return nil
}

func setCount(c *Conn, n int) error {
	// return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, syscall.TCP_KEEPCNT, n))
	return nil
}

func setInterval(c *Conn, secs int) error {
	// return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, syscall.TCP_KEEPINTVL, secs))
	return nil
}
