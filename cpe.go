package lz

type LazyCallPanicOrError0 struct {
	fn Func0Err
}

func CPE0(fn Func0Err) LazyCallPanicOrError0 {
	return LazyCallPanicOrError0{fn}
}

func (lc LazyCallPanicOrError0) Call() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = recoverFromErr(r)
		}
	}()
	return lc.fn()
}

type LazyCallPanicOrError1[T1 any] struct {
	fn Func1Err[T1]
	v1 T1
}

func CPE1[T1 any](fn Func1Err[T1], v1 T1) LazyCallPanicOrError1[T1] {
	return LazyCallPanicOrError1[T1]{fn, v1}
}

func (lc LazyCallPanicOrError1[T1]) Call() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = recoverFromErr(r)
		}
	}()
	return lc.fn(lc.v1)
}

type LazyCallPanicOrError2[T1, T2 any] struct {
	fn Func2Err[T1, T2]
	v1 T1
	v2 T2
}

func CPE2[T1, T2 any](fn Func2Err[T1, T2], v1 T1, v2 T2) LazyCallPanicOrError2[T1, T2] {
	return LazyCallPanicOrError2[T1, T2]{fn, v1, v2}
}

func (lc LazyCallPanicOrError2[T1, T2]) Call() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = recoverFromErr(r)
		}
	}()
	return lc.fn(lc.v1, lc.v2)
}

type LazyCallPanicOrError3[T1 any, T2 any, T3 any] struct {
	fn Func3Err[T1, T2, T3]
	v1 T1
	v2 T2
	v3 T3
}

func CPE3[T1, T2, T3 any](fn Func3Err[T1, T2, T3], v1 T1, v2 T2, v3 T3) LazyCallPanicOrError3[T1, T2, T3] {
	return LazyCallPanicOrError3[T1, T2, T3]{fn, v1, v2, v3}
}

func (lc LazyCallPanicOrError3[T1, T2, T3]) Call() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = recoverFromErr(r)
		}
	}()
	return lc.fn(lc.v1, lc.v2, lc.v3)
}

type LazyCallPanicOrError4[T1, T2, T3, T4 any] struct {
	fn Func4Err[T1, T2, T3, T4]
	v1 T1
	v2 T2
	v3 T3
	v4 T4
}

func CPE4[T1, T2, T3, T4 any](fn Func4Err[T1, T2, T3, T4], v1 T1, v2 T2, v3 T3, v4 T4) LazyCallPanicOrError4[T1, T2, T3, T4] {
	return LazyCallPanicOrError4[T1, T2, T3, T4]{fn, v1, v2, v3, v4}
}

func (lc LazyCallPanicOrError4[T1, T2, T3, T4]) Call() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = recoverFromErr(r)
		}
	}()
	return lc.fn(lc.v1, lc.v2, lc.v3, lc.v4)
}

type LazyCallPanicOrError5[T1, T2, T3, T4, T5 any] struct {
	fn Func5Err[T1, T2, T3, T4, T5]
	v1 T1
	v2 T2
	v3 T3
	v4 T4
	v5 T5
}

func CPE5[T1, T2, T3, T4, T5 any](fn Func5Err[T1, T2, T3, T4, T5], v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) LazyCallPanicOrError5[T1, T2, T3, T4, T5] {
	return LazyCallPanicOrError5[T1, T2, T3, T4, T5]{fn, v1, v2, v3, v4, v5}
}

func (lc LazyCallPanicOrError5[T1, T2, T3, T4, T5]) Call() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = recoverFromErr(r)
		}
	}()
	return lc.fn(lc.v1, lc.v2, lc.v3, lc.v4, lc.v5)
}

type LazyCallPanicOrError6[T1, T2, T3, T4, T5, T6 any] struct {
	fn Func6Err[T1, T2, T3, T4, T5, T6]
	v1 T1
	v2 T2
	v3 T3
	v4 T4
	v5 T5
	v6 T6
}

func CPE6[T1, T2, T3, T4, T5, T6 any](fn Func6Err[T1, T2, T3, T4, T5, T6], v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) LazyCallPanicOrError6[T1, T2, T3, T4, T5, T6] {
	return LazyCallPanicOrError6[T1, T2, T3, T4, T5, T6]{fn, v1, v2, v3, v4, v5, v6}
}

func (lc LazyCallPanicOrError6[T1, T2, T3, T4, T5, T6]) Call() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = recoverFromErr(r)
		}
	}()
	return lc.fn(lc.v1, lc.v2, lc.v3, lc.v4, lc.v5, lc.v6)
}

type LazyCallPanicOrError7[T1, T2, T3, T4, T5, T6, T7 any] struct {
	fn Func7Err[T1, T2, T3, T4, T5, T6, T7]
	v1 T1
	v2 T2
	v3 T3
	v4 T4
	v5 T5
	v6 T6
	v7 T7
}

func CPE7[T1, T2, T3, T4, T5, T6, T7 any](fn Func7Err[T1, T2, T3, T4, T5, T6, T7], v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7) LazyCallPanicOrError7[T1, T2, T3, T4, T5, T6, T7] {
	return LazyCallPanicOrError7[T1, T2, T3, T4, T5, T6, T7]{fn, v1, v2, v3, v4, v5, v6, v7}
}

func (lc LazyCallPanicOrError7[T1, T2, T3, T4, T5, T6, T7]) Call() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = recoverFromErr(r)
		}
	}()
	return lc.fn(lc.v1, lc.v2, lc.v3, lc.v4, lc.v5, lc.v6, lc.v7)
}

type LazyCallPanicOrError8[T1, T2, T3, T4, T5, T6, T7, T8 any] struct {
	fn Func8Err[T1, T2, T3, T4, T5, T6, T7, T8]
	v1 T1
	v2 T2
	v3 T3
	v4 T4
	v5 T5
	v6 T6
	v7 T7
	v8 T8
}

func CPE8[T1, T2, T3, T4, T5, T6, T7, T8 any](fn Func8Err[T1, T2, T3, T4, T5, T6, T7, T8], v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8) LazyCallPanicOrError8[T1, T2, T3, T4, T5, T6, T7, T8] {
	return LazyCallPanicOrError8[T1, T2, T3, T4, T5, T6, T7, T8]{fn, v1, v2, v3, v4, v5, v6, v7, v8}
}

func (lc LazyCallPanicOrError8[T1, T2, T3, T4, T5, T6, T7, T8]) Call() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = recoverFromErr(r)
		}
	}()
	return lc.fn(lc.v1, lc.v2, lc.v3, lc.v4, lc.v5, lc.v6, lc.v7, lc.v8)
}

type LazyCallPanicOrError9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any] struct {
	fn Func9Err[T1, T2, T3, T4, T5, T6, T7, T8, T9]
	v1 T1
	v2 T2
	v3 T3
	v4 T4
	v5 T5
	v6 T6
	v7 T7
	v8 T8
	v9 T9
}

func CPE9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](fn Func9Err[T1, T2, T3, T4, T5, T6, T7, T8, T9], v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9) LazyCallPanicOrError9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	return LazyCallPanicOrError9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{fn, v1, v2, v3, v4, v5, v6, v7, v8, v9}
}

func (lc LazyCallPanicOrError9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Call() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = recoverFromErr(r)
		}
	}()
	return lc.fn(lc.v1, lc.v2, lc.v3, lc.v4, lc.v5, lc.v6, lc.v7, lc.v8, lc.v9)
}
