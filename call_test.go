package lz_test

import (
	"errors"
	"testing"
	"time"

	"github.com/jopbrown/lz"
	"github.com/stretchr/testify/assert"
)

func sleepAndError(d time.Duration, err error) error {
	time.Sleep(d)
	return err
}

func sleepAndPanic(d time.Duration, err error) {
	time.Sleep(d)
	if err != nil {
		panic(err)
	}
}

func sleepAndPanicOrError(d time.Duration, perr, rerr error) error {
	time.Sleep(d)
	if perr != nil {
		panic(perr)
	}

	return rerr
}

var (
	errGolden = errors.New("golden")
	errWrong1 = errors.New("wrong1")
	errWrong2 = errors.New("wrong2")
	errWrong3 = errors.New("wrong3")
	errWrong4 = errors.New("wrong4")
	errWrong5 = errors.New("wrong5")
)

func TestCall(t *testing.T) {
	err := lz.Call(
		lz.CE2(sleepAndError, 1*time.Microsecond, nil),
		lz.CE2(sleepAndError, 2*time.Microsecond, nil),
		lz.CE2(sleepAndError, 2*time.Microsecond, errGolden),
		lz.CE2(sleepAndError, 1*time.Microsecond, errWrong1),
		lz.CE2(sleepAndError, 1*time.Microsecond, errWrong2),
	)

	assert.Equal(t, errGolden, err)

	err = lz.Call(
		lz.CE2(sleepAndError, 1*time.Microsecond, nil),
		lz.CP2(sleepAndPanic, 2*time.Microsecond, nil),
		lz.CE2(sleepAndError, 2*time.Microsecond, nil),
		lz.CE2(sleepAndError, 1*time.Microsecond, nil),
		lz.CE2(sleepAndError, 1*time.Microsecond, nil),
	)

	assert.NoError(t, err)
}

func TestCallParallel(t *testing.T) {
	err := lz.CallInParallel(
		lz.CE2(sleepAndError, 3*time.Microsecond, errWrong1),
		lz.CE2(sleepAndError, 3*time.Microsecond, errWrong2),
		lz.CP2(sleepAndPanic, 1*time.Microsecond, errGolden),
		lz.CE2(sleepAndError, 2*time.Microsecond, errWrong3),
		lz.CE2(sleepAndError, 2*time.Microsecond, errWrong4),
	)

	assert.Error(t, errGolden, err)

	err = lz.CallInParallel(
		lz.CE2(sleepAndError, 3*time.Microsecond, nil),
		lz.CE2(sleepAndError, 3*time.Microsecond, nil),
		lz.CE2(sleepAndError, 1*time.Microsecond, nil),
		lz.CE2(sleepAndError, 2*time.Microsecond, nil),
		lz.CE2(sleepAndError, 2*time.Microsecond, nil),
	)

	assert.NoError(t, err)
}

func TestCallParallelLimit(t *testing.T) {
	err := lz.CallInParallelLimit(1,
		lz.CE2(sleepAndError, 3*time.Microsecond, errGolden),
		lz.CE2(sleepAndError, 3*time.Microsecond, errWrong1),
		lz.CE2(sleepAndError, 1*time.Microsecond, errWrong2),
		lz.CE2(sleepAndError, 2*time.Microsecond, errWrong3),
		lz.CE2(sleepAndError, 2*time.Microsecond, errWrong4),
	)

	assert.Error(t, errGolden, err)

	err = lz.CallInParallelLimit(1,
		lz.CE2(sleepAndError, 3*time.Microsecond, nil),
		lz.CE2(sleepAndError, 3*time.Microsecond, nil),
		lz.CE2(sleepAndError, 1*time.Microsecond, nil),
		lz.CE2(sleepAndError, 2*time.Microsecond, nil),
		lz.CE2(sleepAndError, 2*time.Microsecond, nil),
	)

	assert.NoError(t, err)
}

func TestCallAll(t *testing.T) {
	err := lz.CallAll(
		lz.CE2(sleepAndError, 1*time.Microsecond, nil),
		lz.CE2(sleepAndError, 2*time.Microsecond, nil),
		lz.CPE3(sleepAndPanicOrError, 2*time.Microsecond, errGolden, nil),
		lz.CP2(sleepAndPanic, 1*time.Microsecond, errWrong1),
		lz.CP2(sleepAndPanic, 1*time.Microsecond, errWrong2),
	)

	assert.True(t, errors.Is(err, errGolden))
	assert.True(t, errors.Is(err, errWrong1))
	assert.True(t, errors.Is(err, errWrong2))
	assert.False(t, errors.Is(err, errWrong3))
	assert.False(t, errors.Is(err, errWrong4))
	assert.False(t, errors.Is(err, errWrong5))

	err = lz.CallAll(
		lz.CE2(sleepAndError, 1*time.Microsecond, nil),
		lz.CE2(sleepAndError, 2*time.Microsecond, nil),
		lz.CE2(sleepAndError, 2*time.Microsecond, nil),
		lz.CE2(sleepAndError, 1*time.Microsecond, nil),
		lz.CE2(sleepAndError, 1*time.Microsecond, nil),
	)

	assert.NoError(t, err)
}

func TestCallAllParallel(t *testing.T) {
	err := lz.CallAllInParallel(
		lz.CE2(sleepAndError, 3*time.Microsecond, errWrong1),
		lz.CE2(sleepAndError, 3*time.Microsecond, errWrong2),
		lz.CE2(sleepAndError, 1*time.Microsecond, errGolden),
		lz.CE2(sleepAndError, 2*time.Microsecond, errWrong3),
		lz.CE2(sleepAndError, 2*time.Microsecond, errWrong4),
	)

	assert.True(t, errors.Is(err, errGolden))
	assert.True(t, errors.Is(err, errWrong1))
	assert.True(t, errors.Is(err, errWrong2))
	assert.True(t, errors.Is(err, errWrong3))
	assert.True(t, errors.Is(err, errWrong4))
	assert.False(t, errors.Is(err, errWrong5))

	err = lz.CallAllInParallel(
		lz.CE2(sleepAndError, 3*time.Microsecond, nil),
		lz.CE2(sleepAndError, 3*time.Microsecond, nil),
		lz.CE2(sleepAndError, 1*time.Microsecond, nil),
		lz.CE2(sleepAndError, 2*time.Microsecond, nil),
		lz.CE2(sleepAndError, 2*time.Microsecond, nil),
	)

	t.Log(err)

	assert.NoError(t, err)
}

func TestCallAllParallelLimit(t *testing.T) {
	err := lz.CallAllInParallelLimit(1,
		lz.CE2(sleepAndError, 3*time.Microsecond, errGolden),
		lz.CE2(sleepAndError, 3*time.Microsecond, errWrong1),
		lz.CE2(sleepAndError, 1*time.Microsecond, errWrong2),
		lz.CE2(sleepAndError, 2*time.Microsecond, errWrong3),
		lz.CE2(sleepAndError, 2*time.Microsecond, errWrong4),
	)

	assert.True(t, errors.Is(err, errGolden))
	assert.True(t, errors.Is(err, errWrong1))
	assert.True(t, errors.Is(err, errWrong2))
	assert.True(t, errors.Is(err, errWrong3))
	assert.True(t, errors.Is(err, errWrong4))
	assert.False(t, errors.Is(err, errWrong5))

	err = lz.CallAllInParallelLimit(1,
		lz.CE2(sleepAndError, 3*time.Microsecond, nil),
		lz.CE2(sleepAndError, 3*time.Microsecond, nil),
		lz.CE2(sleepAndError, 1*time.Microsecond, nil),
		lz.CE2(sleepAndError, 2*time.Microsecond, nil),
		lz.CE2(sleepAndError, 2*time.Microsecond, nil),
	)

	assert.NoError(t, err)
}
