package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		c chan os.Signal
		s os.Signal
	)

	c = make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM,
		syscall.SIGINT, syscall.SIGSTOP)
	for {
		s = <-c
		switch s {
		case syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP:
			fmt.Println("s=", s)
			return
		default:
			return
		}
	}
	return
}
