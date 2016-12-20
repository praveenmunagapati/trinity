package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/clownpriest/trinity/core/system"
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
		subsys.Process.Process.Signal(sig)
	}
	fmt.Println("\nshutting down core....")
	os.Exit(0)
}

func interruptSubsystem(subsys system.SubsystemType, sig os.Signal) {
	node.Subsystems[subsys].Process.Process.Signal(sig)
}
