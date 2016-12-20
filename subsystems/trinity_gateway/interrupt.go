package main

import (
	"os"
	"os/signal"
	"sync"
)

func startInterruptHandler(sigChan chan os.Signal, wg *sync.WaitGroup) {
	defer wg.Done()
	signal.Notify(sigChan, os.Interrupt, os.Kill)
	for sig := range sigChan {
		shutdown(sig)
	}
}

func shutdown(sig os.Signal) {
	os.Exit(0)
}
