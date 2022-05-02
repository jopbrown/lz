package lz

type Func0Err func() error
type Func1Err[T1 any] func(v T1) error
type Func2Err[T1, T2 any] func(v1 T1, v2 T2) error
type Func3Err[T1, T2, T3 any] func(v1 T1, v2 T2, v3 T3) error
type Func4Err[T1, T2, T3, T4 any] func(v1 T1, v2 T2, v3 T3, v4 T4) error
type Func5Err[T1, T2, T3, T4, T5 any] func(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) error
type Func6Err[T1, T2, T3, T4, T5, T6 any] func(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) error
type Func7Err[T1, T2, T3, T4, T5, T6, T7 any] func(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7) error
type Func8Err[T1, T2, T3, T4, T5, T6, T7, T8 any] func(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8) error
type Func9Err[T1, T2, T3, T4, T5, T6, T7, T8, T9 any] func(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9) error

type Func0 func()
type Func1[T1 any] func(v T1)
type Func2[T1, T2 any] func(v1 T1, v2 T2)
type Func3[T1, T2, T3 any] func(v1 T1, v2 T2, v3 T3)
type Func4[T1, T2, T3, T4 any] func(v1 T1, v2 T2, v3 T3, v4 T4)
type Func5[T1, T2, T3, T4, T5 any] func(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5)
type Func6[T1, T2, T3, T4, T5, T6 any] func(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6)
type Func7[T1, T2, T3, T4, T5, T6, T7 any] func(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7)
type Func8[T1, T2, T3, T4, T5, T6, T7, T8 any] func(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8)
type Func9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any] func(v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9)
