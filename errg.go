package lz

import (
	"sync"
)

type errGroup struct {
	goroutines int
	limitCh    chan struct{}
	mu         sync.Mutex
	stop       chan struct{}
	err        error
	wg         sync.WaitGroup
}

func newErrGroup(goroutines int) *errGroup {
	g := &errGroup{}
	g.goroutines = goroutines
	if goroutines > 0 {
		g.limitCh = make(chan struct{}, goroutines)
	}
	g.stop = make(chan struct{})
	return g
}

func (g *errGroup) Go(f func() error) {
	if g.Err() != nil {
		return
	}

	g.wg.Add(1)
	go func() {
		defer g.wg.Done()

		if g.goroutines > 0 {
			select {
			case g.limitCh <- struct{}{}:
				defer func() { <-g.limitCh }()
			case <-g.stop:
				return
			}
		} else {
			select {
			case <-g.stop:
				return
			default:
			}
		}

		if err := f(); err != nil {
			g.mu.Lock()
			if g.err == nil {
				g.err = err
				close(g.stop)
			}
			g.mu.Unlock()
		}
	}()
}

func (g *errGroup) Err() error {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.err
}

func (g *errGroup) Wait() error {
	g.wg.Wait()
	return g.Err()
}
