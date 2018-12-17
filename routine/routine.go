package routine

import "sync"

// WaitGroup is syntax suger for sync.WaitGroup
type WaitGroup struct {
	Wg sync.WaitGroup
}

// Go run routine with wait-group count
func (wg *WaitGroup) Go(fn func(...interface{}), arg ...interface{}) {
	wg.Wg.Add(1)
	go func() {
		fn(arg...)
		wg.Wg.Done()
	}()
}

// Wait wait utill all routine dead
func (wg *WaitGroup) Wait() {
	wg.Wg.Wait()
}
