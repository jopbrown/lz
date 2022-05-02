package lz

type LazyCall0 struct {
	fn Func0Err
}

func CE0(fn Func0Err) LazyCall0 {
	return LazyCall0{fn}
}

func (lc LazyCall0) Call() error {
	return lc.fn()
}

type LazyCall1[T1 any] struct {
	fn Func1Err[T1]
	v1 T1
}

func CE1[T1 any](fn Func1Err[T1], v1 T1) LazyCall1[T1] {
	return LazyCall1[T1]{fn, v1}
}

func (lc LazyCall1[T1]) Call() error {
	return lc.fn(lc.v1)
}

type LazyCall2[T1, T2 any] struct {
	fn Func2Err[T1, T2]
	v1 T1
	v2 T2
}

func CE2[T1, T2 any](fn Func2Err[T1, T2], v1 T1, v2 T2) LazyCall2[T1, T2] {
	return LazyCall2[T1, T2]{fn, v1, v2}
}

func (lc LazyCall2[T1, T2]) Call() error {
	return lc.fn(lc.v1, lc.v2)
}

type LazyCall3[T1 any, T2 any, T3 any] struct {
	fn Func3Err[T1, T2, T3]
	v1 T1
	v2 T2
	v3 T3
}

func CE3[T1, T2, T3 any](fn Func3Err[T1, T2, T3], v1 T1, v2 T2, v3 T3) LazyCall3[T1, T2, T3] {
	return LazyCall3[T1, T2, T3]{fn, v1, v2, v3}
}

func (lc LazyCall3[T1, T2, T3]) Call() error {
	return lc.fn(lc.v1, lc.v2, lc.v3)
}

type LazyCall4[T1, T2, T3, T4 any] struct {
	fn Func4Err[T1, T2, T3, T4]
	v1 T1
	v2 T2
	v3 T3
	v4 T4
}

func CE4[T1, T2, T3, T4 any](fn Func4Err[T1, T2, T3, T4], v1 T1, v2 T2, v3 T3, v4 T4) LazyCall4[T1, T2, T3, T4] {
	return LazyCall4[T1, T2, T3, T4]{fn, v1, v2, v3, v4}
}

func (lc LazyCall4[T1, T2, T3, T4]) Call() error {
	return lc.fn(lc.v1, lc.v2, lc.v3, lc.v4)
}

type LazyCall5[T1, T2, T3, T4, T5 any] struct {
	fn Func5Err[T1, T2, T3, T4, T5]
	v1 T1
	v2 T2
	v3 T3
	v4 T4
	v5 T5
}

func CE5[T1, T2, T3, T4, T5 any](fn Func5Err[T1, T2, T3, T4, T5], v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) LazyCall5[T1, T2, T3, T4, T5] {
	return LazyCall5[T1, T2, T3, T4, T5]{fn, v1, v2, v3, v4, v5}
}

func (lc LazyCall5[T1, T2, T3, T4, T5]) Call() error {
	return lc.fn(lc.v1, lc.v2, lc.v3, lc.v4, lc.v5)
}

type LazyCall6[T1, T2, T3, T4, T5, T6 any] struct {
	fn Func6Err[T1, T2, T3, T4, T5, T6]
	v1 T1
	v2 T2
	v3 T3
	v4 T4
	v5 T5
	v6 T6
}

func CE6[T1, T2, T3, T4, T5, T6 any](fn Func6Err[T1, T2, T3, T4, T5, T6], v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) LazyCall6[T1, T2, T3, T4, T5, T6] {
	return LazyCall6[T1, T2, T3, T4, T5, T6]{fn, v1, v2, v3, v4, v5, v6}
}

func (lc LazyCall6[T1, T2, T3, T4, T5, T6]) Call() error {
	return lc.fn(lc.v1, lc.v2, lc.v3, lc.v4, lc.v5, lc.v6)
}

type LazyCall7[T1, T2, T3, T4, T5, T6, T7 any] struct {
	fn Func7Err[T1, T2, T3, T4, T5, T6, T7]
	v1 T1
	v2 T2
	v3 T3
	v4 T4
	v5 T5
	v6 T6
	v7 T7
}

func CE7[T1, T2, T3, T4, T5, T6, T7 any](fn Func7Err[T1, T2, T3, T4, T5, T6, T7], v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7) LazyCall7[T1, T2, T3, T4, T5, T6, T7] {
	return LazyCall7[T1, T2, T3, T4, T5, T6, T7]{fn, v1, v2, v3, v4, v5, v6, v7}
}

func (lc LazyCall7[T1, T2, T3, T4, T5, T6, T7]) Call() error {
	return lc.fn(lc.v1, lc.v2, lc.v3, lc.v4, lc.v5, lc.v6, lc.v7)
}

type LazyCall8[T1, T2, T3, T4, T5, T6, T7, T8 any] struct {
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

func CE8[T1, T2, T3, T4, T5, T6, T7, T8 any](fn Func8Err[T1, T2, T3, T4, T5, T6, T7, T8], v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8) LazyCall8[T1, T2, T3, T4, T5, T6, T7, T8] {
	return LazyCall8[T1, T2, T3, T4, T5, T6, T7, T8]{fn, v1, v2, v3, v4, v5, v6, v7, v8}
}

func (lc LazyCall8[T1, T2, T3, T4, T5, T6, T7, T8]) Call() error {
	return lc.fn(lc.v1, lc.v2, lc.v3, lc.v4, lc.v5, lc.v6, lc.v7, lc.v8)
}

type LazyCall9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any] struct {
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

func CE9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](fn Func9Err[T1, T2, T3, T4, T5, T6, T7, T8, T9], v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9) LazyCall9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	return LazyCall9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{fn, v1, v2, v3, v4, v5, v6, v7, v8, v9}
}

func (lc LazyCall9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Call() error {
	return lc.fn(lc.v1, lc.v2, lc.v3, lc.v4, lc.v5, lc.v6, lc.v7, lc.v8, lc.v9)
}
