package routine

import "sync"

type WaitGroup struct {
	Wg sync.WaitGroup
}

func (wg *WaitGroup) Go(fn func(...interface{}), arg ...interface{}) {
	wg.Wg.Add(1)
	go func() {
		fn(arg)
		wg.Wg.Done()
	}()
}

func (wg *WaitGroup) Wait() {
	wg.Wg.Wait()
}
