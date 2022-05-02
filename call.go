package lz

import (
	"github.com/hashicorp/go-multierror"
)

type Caller interface {
	Call() error
}

// Call multiple lazy objects
func Call(cs ...Caller) error {
	var err error
	for _, c := range cs {
		if err = c.Call(); err != nil {
			return err
		}
	}
	return nil
}

// CallInParallel call multiple lazy objects in parallel and stop on the first occur error.
func CallInParallel(cs ...Caller) error {
	return CallInParallelLimit(0, cs...)
}

// CallInParallelLimit call multiple lazy objects in parallel with limited goroutines
// and stop on the first occur error.
func CallInParallelLimit(goroutines int, cs ...Caller) error {
	eg := newErrGroup(goroutines)
	for _, c := range cs {
		c := c
		eg.Go(func() error {
			return c.Call()
		})
	}
	return eg.Wait()
}

// CallAll call all lazy objects in order and return multierror if any.
func CallAll(cs ...Caller) error {
	var merr error
	for _, c := range cs {
		if err := c.Call(); err != nil {
			merr = multierror.Append(merr, err)
		}
	}
	return merr
}

// CallAllInParallel call all lazy objects in parallel and return multierror if any.
func CallAllInParallel(cs ...Caller) error {
	eg := &multierror.Group{}
	for _, c := range cs {
		eg.Go(c.Call)
	}
	return eg.Wait().ErrorOrNil()
}

// CallAllInParallelLimit call all lazy objects in parallel with limited goroutines and return multierror if any.
func CallAllInParallelLimit(goroutines int, cs ...Caller) error {
	eg := &multierror.Group{}
	ch := make(chan struct{}, goroutines)
	defer close(ch)
	for _, c := range cs {
		c := c
		eg.Go(func() error {
			ch <- struct{}{}
			defer func() { <-ch }()
			return c.Call()
		})
	}
	return eg.Wait().ErrorOrNil()
}
