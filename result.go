package result

type Result[T any] struct {
	val T
	err error
}

func NewResult[T any](val T, err error) Result[T] {
	return Result[T]{val, err}
}

func Ok[T any](val T) Result[T] {
	return Result[T]{val, nil}
}

func Err[T any](err error) Result[T] {
	return Result[T]{err: err}
}

func (r Result[T]) AsTuple() (val T, err error) {
	return r.val, r.err
}

func (r Result[T]) UnwrapOr(defaultValue T) T {
	if r.err != nil {
		return defaultValue
	}
	return r.val
}

func (r Result[T]) Unwrap() T {
	if r.err != nil {
		panic(r.err)
	}
	return r.val
}

func (r Result[T]) Expect(msg string) T {
	if r.err != nil {
		panic(msg)
	}
	return r.val
}

func (r Result[T]) UnwrapOrElse(f func() T) T {
	if r.err != nil {
		return f()
	}
	return r.val
}

func (r Result[T]) UnwrapOrDefault() T {
	return r.val
}

func (r Result[T]) IsErr() bool {
	return !r.IsOk()
}

func (r Result[T]) IsOk() bool {
	return r.err == nil
}
