package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func startInterruptHandler(sigChan chan os.Signal, wg *sync.WaitGroup) {
	defer wg.Done()
	signal.Notify(sigChan, os.Interrupt, os.Kill)
	for sig := range sigChan {
		cleanupSubsystems(sig)
	}
}

func cleanupSubsystems(sig os.Signal) {
	fmt.Printf("\n\nrecieved an interrupt signal....shutting down\n\n")
	for key, subsys := range node.Subsystems {
		fmt.Printf("subsystem shutdown:\t%s\n", key)
		subsys.Process.Process.Kill()
	}
	fmt.Println("\nshutting down core....")
	os.Exit(0)

}
