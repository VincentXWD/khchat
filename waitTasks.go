package main

import (
	log "github.com/Sirupsen/logrus"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	lock    sync.RWMutex
	onceTer sync.Once
)

func Task(f func()) {
	lock.RLock()
	defer lock.RUnlock()

	f()
}

func Terminate() {
	log.Println("WAITING FOR TASK FINISHED")
	lock.Lock()
	defer lock.Unlock()
	os.Exit(0)
}

func init() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		_ = <-c
		onceTer.Do(Terminate)
	}()
}
