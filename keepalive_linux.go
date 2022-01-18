package tcpkeepalive

import (
	"os"
	"syscall"
)

//https://stackoverflow.com/questions/5907527/application-control-of-tcp-retransmission-on-linux
func setTcpUsertimeout(fd int, secs int) error {

	return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, 18, secs*1000))
}
func setIdle(fd int, secs int) error {
	return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, syscall.TCP_KEEPIDLE, secs))
}

func setCount(fd int, n int) error {
	return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, syscall.TCP_KEEPCNT, n))
}

func setInterval(fd int, secs int) error {
	return os.NewSyscallError("setsockopt", syscall.SetsockoptInt(fd, syscall.IPPROTO_TCP, syscall.TCP_KEEPINTVL, secs))
}
