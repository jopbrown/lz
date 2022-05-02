# lz

**`lz` is a Lazy Call Go library based on Go 1.18+ Generics.**

You can define a function and parameters and **call-by-need** later.

## Example

```go
import (
    "github.com/jopbrown/lz"
)

func Download(url string) error {
	// download file from url
}

lazyCall1 := lz.CE1(Download, "http://example.com/mirror1/file.zip")
lazyCall2 := lz.CE1(Download, "http://example.com/mirror2/file.zip")
lazyCall3 := lz.CE1(Download, "http://example.com/mirror3/file.zip")

// call in order
err := lz.Call(lazyCall1, lazyCall2, lazyCall3)
if err != nil {
	// handle  error
}

// call in parallel
err = lz.CallInParallel(lazyCall1, lazyCall2, lazyCall3)
if err != nil {
	// handle  error
}

```

## Usage

### Create lazy call object

- `lz.CE{0~9}`

	Create a lazy call object with a function return an `error`.

	```go
	func DoSomething() error {
		...
	}
	func DoSomethingWith1Param(v1 string) error {
		...
	}
	func DoSomethingWith2Param(v1 int, v2 float64) error {
		...
	}

	lzc1 := lz.CE0(DoSomething)
	lzc2 := lz.CE1(DoSomethingWith1Param, v1)
	lzc3 := lz.CE2(DoSomethingWith2Param, v1, v2)
	```

- `lz.CP{0~9}`

	Create a lazy call object with a function has no return but may panic.

	When called later, panic will be catched and return as `error`.

	```go
	func DoSomething() {
		...
		panic("something wrong")
	}
	func DoSomethingWith1Param(v1 string) {
		...
		panic("something wrong")
	}
	func DoSomethingWith2Param(v1 int, v2 float64) {
		...
		panic("something wrong")
	}

	lzc1 := lz.CP0(DoSomething)
	lzc2 := lz.CP1(DoSomethingWith1Param, v1)
	lzc3 := lz.CP2(DoSomethingWith2Param, v1, v2)
	```

- `lz.CEP{0~9}`

	Create a lazy call object with a function return an `error` and may panic.

	When called later, panic will be catched and return as `error`.

	```go
	func DoSomething() error {
		...
		panic("something wrong")
	}
	func DoSomethingWith1Param(v1 string) error {
		...
		panic("something wrong")
	}
	func DoSomethingWith2Param(v1 int, v2 float64) error {
		...
		panic("something wrong")
	}


	lzc1 := lz.CPE0(DoSomething)
	lzc2 := lz.CPE1(DoSomethingWith1Param, v1)
	lzc3 := lz.CPE2(DoSomethingWith2Param, v1, v2)
	```

### Call single lazy object directly

```go
var err := lzc1.Call()
if err != nil {
	// handle  error
}
```

### Call multiple lazy objects

- Call in order and stop on the first occur error.

	```go
	func sleepAndError(d time.Duration, err error) error {
		time.Sleep(d)
		return err
	}

	err := lz.Call(
		lz.CE2(sleepAndError, 1*time.Microsecond, nil),
		lz.CE2(sleepAndError, 2*time.Microsecond, nil),
		lz.CE2(sleepAndError, 2*time.Microsecond, errGolden),
		lz.CE2(sleepAndError, 1*time.Microsecond, errWrong1),
		lz.CE2(sleepAndError, 1*time.Microsecond, errWrong2),
	)

	errors.Is(err, errGolden) // true
	```

- Call in parallel and stop on the first occur error.

	```go
	err := lz.CallInParallel(
		lz.CE2(sleepAndError, 3*time.Microsecond, errWrong1),
		lz.CP2(sleepAndPanic, 3*time.Microsecond, errWrong2),
		lz.CP2(sleepAndPanic, 1*time.Microsecond, errWrong3),
		lz.CE2(sleepAndError, 2*time.Microsecond, errWrong4),
		lz.CE2(sleepAndError, 2*time.Microsecond, errWrong5),
	)

	// which error is unpredictable
	if err != nil {
		// handle  error
	}
	```

- Call in parallel with limited goroutines and stop on the first occur error.

	```go
	err := lz.CallInParallelLimit(2, // limit to 2 goroutines
			lz.CE2(sleepAndError, 3*time.Microsecond, errGolden),
			lz.CE2(sleepAndError, 3*time.Microsecond, errWrong1),
			lz.CE2(sleepAndError, 1*time.Microsecond, errWrong2),
			lz.CE2(sleepAndError, 2*time.Microsecond, errWrong3),
			lz.CE2(sleepAndError, 2*time.Microsecond, errWrong4),
		)

	// which error is unpredictable
	if err != nil {
		// handle  error
	}
	```

- Call all lazy objects in order and return [`multierror`](https://pkg.go.dev/github.com/hashicorp/go-multierror@v1.1.1#Error) if any.

	```go
	err := lz.CallAll(
			lz.CE2(sleepAndError, 1*time.Microsecond, errWrong1),
			lz.CE2(sleepAndError, 2*time.Microsecond, nil),
			lz.CE3(sleepAndError, 2*time.Microsecond, nil),
			lz.CP2(sleepAndPanic, 1*time.Microsecond, errWrong2),
			lz.CP2(sleepAndPanic, 1*time.Microsecond, errWrong3),
		)

	errors.Is(err, errWrong1) // true
	errors.Is(err, errWrong2) // true
	errors.Is(err, errWrong3) // true
	```

- Call all lazy objects in parallel and return [`multierror`](https://pkg.go.dev/github.com/hashicorp/go-multierror@v1.1.1#Error) if any.

	```go
	err := lz.CallAllInParallel(
			lz.CE2(sleepAndError, 1*time.Microsecond, errWrong1),
			lz.CE2(sleepAndError, 2*time.Microsecond, nil),
			lz.CE3(sleepAndError, 2*time.Microsecond, nil),
			lz.CP2(sleepAndPanic, 1*time.Microsecond, errWrong2),
			lz.CP2(sleepAndPanic, 1*time.Microsecond, errWrong3),
		)

	errors.Is(err, errWrong1) // true
	errors.Is(err, errWrong2) // true
	errors.Is(err, errWrong3) // true
	```

- Call all lazy objects in parallel with limited goroutines and return [`multierror`](https://pkg.go.dev/github.com/hashicorp/go-multierror@v1.1.1#Error) if any.

	```go
	err := lz.CallAllInParallelLimit(2, // limit to 2 goroutines
			lz.CE2(sleepAndError, 1*time.Microsecond, errWrong1),
			lz.CE2(sleepAndError, 2*time.Microsecond, nil),
			lz.CE3(sleepAndError, 2*time.Microsecond, nil),
			lz.CP2(sleepAndPanic, 1*time.Microsecond, errWrong2),
			lz.CP2(sleepAndPanic, 1*time.Microsecond, errWrong3),
		)

	errors.Is(err, errWrong1) // true
	errors.Is(err, errWrong2) // true
	errors.Is(err, errWrong3) // true
	```
