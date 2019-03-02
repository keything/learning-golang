package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// StartSignal register signals handler.
func main() {
	var (
		c chan os.Signal
		s os.Signal
	)
	c = make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM,
		syscall.SIGINT, syscall.SIGSTOP, syscall.SIGSEGV)
	// Block until a signal is received.
	for {
		s = <-c
		fmt.Println("get a signal ", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT, syscall.SIGSEGV:
			return
		case syscall.SIGHUP:
			// TODO reload
			//return
		default:
			return
		}
	}
}
