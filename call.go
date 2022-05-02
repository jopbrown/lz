package lz

import (
	"github.com/hashicorp/go-multierror"
	"golang.org/x/sync/errgroup"
)

type Caller interface {
	Call() error
}

func Call(cs ...Caller) error {
	var err error
	for _, c := range cs {
		if err = c.Call(); err != nil {
			return err
		}
	}
	return nil
}

func CallInParallel(cs ...Caller) error {
	eg := &errgroup.Group{}
	for _, c := range cs {
		eg.Go(c.Call)
	}
	return eg.Wait()
}

func CallInParallelLimit(goroutines int, cs ...Caller) error {
	eg := &errgroup.Group{}
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
	return eg.Wait()
}

func CallAll(cs ...Caller) error {
	var merr error
	for _, c := range cs {
		if err := c.Call(); err != nil {
			merr = multierror.Append(merr, err)
		}
	}
	return merr
}

func CallAllInParallel(cs ...Caller) error {
	eg := &multierror.Group{}
	for _, c := range cs {
		eg.Go(c.Call)
	}
	return eg.Wait().ErrorOrNil()
}

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
