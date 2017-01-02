package system

import (
	"os"
	"os/signal"
	"sync"
)

type InterruptHandler struct {
	sig chan os.Signal
	wg  sync.WaitGroup
}

/*
NewInteruptHandler returns a pointer to a new InteruptHandler
struct
*/
func NewInterruptHandler() *InterruptHandler {
	ih := &InterruptHandler{}
	ih.sig = make(chan os.Signal, 1)
	return ih
}

// Handle starts handling the given signals, and will call the handler
// callback function each time a signal is catched. The function is passed
// the number of times the handler has been triggered in total, as
// well as the handler itself, so that the handling logic can use the
// handler's wait group to ensure clean shutdown when Close() is called.
func (ih *InterruptHandler) Handle(handler func(count int, ih *InterruptHandler), sigs ...os.Signal) {
	signal.Notify(ih.sig, sigs...)
	ih.wg.Add(1)
	go func() {
		defer ih.wg.Done()
		count := 0
		for _ = range ih.sig {
			count++
			handler(count, ih)
		}
		signal.Stop(ih.sig)
	}()
}

/*
Close the interupt handler's channels
*/
func (intHandle *InterruptHandler) Close() error {
	close(intHandle.sig)
	intHandle.wg.Wait()
	return nil
}
