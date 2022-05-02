package lz

type LazyCallPanic0 struct {
	fn Func0
}

func CP0(fn Func0) LazyCallPanic0 {
	return LazyCallPanic0{fn}
}

func (lc LazyCallPanic0) Call() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = recoverFromErr(r)
		}
	}()
	lc.fn()
	return nil
}

type LazyCallPanic1[T1 any] struct {
	fn Func1[T1]
	v1 T1
}

func CP1[T1 any](fn Func1[T1], v1 T1) LazyCallPanic1[T1] {
	return LazyCallPanic1[T1]{fn, v1}
}

func (lc LazyCallPanic1[T1]) Call() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = recoverFromErr(r)
		}
	}()
	lc.fn(lc.v1)
	return nil
}

type LazyCallPanic2[T1, T2 any] struct {
	fn Func2[T1, T2]
	v1 T1
	v2 T2
}

func CP2[T1, T2 any](fn Func2[T1, T2], v1 T1, v2 T2) LazyCallPanic2[T1, T2] {
	return LazyCallPanic2[T1, T2]{fn, v1, v2}
}

func (lc LazyCallPanic2[T1, T2]) Call() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = recoverFromErr(r)
		}
	}()
	lc.fn(lc.v1, lc.v2)
	return nil
}

type LazyCallPanic3[T1, T2, T3 any] struct {
	fn Func3[T1, T2, T3]
	v1 T1
	v2 T2
	v3 T3
}

func CP3[T1, T2, T3 any](fn Func3[T1, T2, T3], v1 T1, v2 T2, v3 T3) LazyCallPanic3[T1, T2, T3] {
	return LazyCallPanic3[T1, T2, T3]{fn, v1, v2, v3}
}

func (lc LazyCallPanic3[T1, T2, T3]) Call() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = recoverFromErr(r)
		}
	}()
	lc.fn(lc.v1, lc.v2, lc.v3)
	return nil
}

type LazyCallPanic4[T1, T2, T3, T4 any] struct {
	fn Func4[T1, T2, T3, T4]
	v1 T1
	v2 T2
	v3 T3
	v4 T4
}

func CP4[T1, T2, T3, T4 any](fn Func4[T1, T2, T3, T4], v1 T1, v2 T2, v3 T3, v4 T4) LazyCallPanic4[T1, T2, T3, T4] {
	return LazyCallPanic4[T1, T2, T3, T4]{fn, v1, v2, v3, v4}
}

func (lc LazyCallPanic4[T1, T2, T3, T4]) Call() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = recoverFromErr(r)
		}
	}()
	lc.fn(lc.v1, lc.v2, lc.v3, lc.v4)
	return nil
}

type LazyCallPanic5[T1, T2, T3, T4, T5 any] struct {
	fn Func5[T1, T2, T3, T4, T5]
	v1 T1
	v2 T2
	v3 T3
	v4 T4
	v5 T5
}

func CP5[T1, T2, T3, T4, T5 any](fn Func5[T1, T2, T3, T4, T5], v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) LazyCallPanic5[T1, T2, T3, T4, T5] {
	return LazyCallPanic5[T1, T2, T3, T4, T5]{fn, v1, v2, v3, v4, v5}
}

func (lc LazyCallPanic5[T1, T2, T3, T4, T5]) Call() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = recoverFromErr(r)
		}
	}()
	lc.fn(lc.v1, lc.v2, lc.v3, lc.v4, lc.v5)
	return nil
}

type LazyCallPanic6[T1, T2, T3, T4, T5, T6 any] struct {
	fn Func6[T1, T2, T3, T4, T5, T6]
	v1 T1
	v2 T2
	v3 T3
	v4 T4
	v5 T5
	v6 T6
}

func CP6[T1, T2, T3, T4, T5, T6 any](fn Func6[T1, T2, T3, T4, T5, T6], v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) LazyCallPanic6[T1, T2, T3, T4, T5, T6] {
	return LazyCallPanic6[T1, T2, T3, T4, T5, T6]{fn, v1, v2, v3, v4, v5, v6}
}

func (lc LazyCallPanic6[T1, T2, T3, T4, T5, T6]) Call() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = recoverFromErr(r)
		}
	}()
	lc.fn(lc.v1, lc.v2, lc.v3, lc.v4, lc.v5, lc.v6)
	return nil
}

type LazyCallPanic7[T1, T2, T3, T4, T5, T6, T7 any] struct {
	fn Func7[T1, T2, T3, T4, T5, T6, T7]
	v1 T1
	v2 T2
	v3 T3
	v4 T4
	v5 T5
	v6 T6
	v7 T7
}

func CP7[T1, T2, T3, T4, T5, T6, T7 any](fn Func7[T1, T2, T3, T4, T5, T6, T7], v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7) LazyCallPanic7[T1, T2, T3, T4, T5, T6, T7] {
	return LazyCallPanic7[T1, T2, T3, T4, T5, T6, T7]{fn, v1, v2, v3, v4, v5, v6, v7}
}

func (lc LazyCallPanic7[T1, T2, T3, T4, T5, T6, T7]) Call() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = recoverFromErr(r)
		}
	}()
	lc.fn(lc.v1, lc.v2, lc.v3, lc.v4, lc.v5, lc.v6, lc.v7)
	return nil
}

type LazyCallPanic8[T1, T2, T3, T4, T5, T6, T7, T8 any] struct {
	fn Func8[T1, T2, T3, T4, T5, T6, T7, T8]
	v1 T1
	v2 T2
	v3 T3
	v4 T4
	v5 T5
	v6 T6
	v7 T7
	v8 T8
}

func CP8[T1, T2, T3, T4, T5, T6, T7, T8 any](fn Func8[T1, T2, T3, T4, T5, T6, T7, T8], v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8) LazyCallPanic8[T1, T2, T3, T4, T5, T6, T7, T8] {
	return LazyCallPanic8[T1, T2, T3, T4, T5, T6, T7, T8]{fn, v1, v2, v3, v4, v5, v6, v7, v8}
}

func (lc LazyCallPanic8[T1, T2, T3, T4, T5, T6, T7, T8]) Call() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = recoverFromErr(r)
		}
	}()
	lc.fn(lc.v1, lc.v2, lc.v3, lc.v4, lc.v5, lc.v6, lc.v7, lc.v8)
	return nil
}

type LazyCallPanic9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any] struct {
	fn Func9[T1, T2, T3, T4, T5, T6, T7, T8, T9]
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

func CP9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](fn Func9[T1, T2, T3, T4, T5, T6, T7, T8, T9], v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9) LazyCallPanic9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	return LazyCallPanic9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{fn, v1, v2, v3, v4, v5, v6, v7, v8, v9}
}

func (lc LazyCallPanic9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Call() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = recoverFromErr(r)
		}
	}()
	lc.fn(lc.v1, lc.v2, lc.v3, lc.v4, lc.v5, lc.v6, lc.v7, lc.v8, lc.v9)
	return nil
}
