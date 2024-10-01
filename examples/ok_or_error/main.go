package main

import (
	"fmt"

	"github.com/olegbannyi/result"
)

type Number interface {
	int | float64
}

type MathOpError int

const (
	DivideZeroError MathOpError = 1 << iota
	SomeOtherError
)

func (e MathOpError) Error() string {
	switch e {
	case DivideZeroError:
		return "cannot be divided by zero"
	default:
		return "some math op error"
	}
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panic on Expect error: %s\n", err)
		}
	}()

	isOk := divideInt(6, 2).IsOk()
	isErr := divideInt(6, 2).IsErr()

	fmt.Printf("divideInt(6, 2) => isOk: %v, isErr: %v\n", isOk, isErr)

	isOk = divideInt(6, 0).IsOk()
	isErr = divideInt(6, 0).IsErr()

	fmt.Printf("divideInt(6, 0) => isOk: %v, isErr: %v\n", isOk, isErr)

	isOk = divideGeneric(10, 2).IsOk()
	isErr = divideGeneric(6, 2).IsErr()

	fmt.Printf("divideGeneric(10, 2) => isOk: %v, isErr: %v\n", isOk, isErr)

	isOk = divideGeneric(3.6, 2).IsOk()
	isErr = divideGeneric(3.6, 2).IsErr()

	fmt.Printf("divideGeneric(3.6, 2) => isOk: %v, isErr: %v\n", isOk, isErr)

	isOk = divideGeneric(3.6, 0).IsOk()
	isErr = divideGeneric(3.6, 0).IsErr()

	fmt.Printf("divideGeneric(3.6, 0) => isOk: %v, isErr: %v\n", isOk, isErr)
}

func divideInt(x, y int) result.Result[int] {
	if y == 0 {
		return result.Err[int](DivideZeroError)
	}

	res := x / y
	return result.Ok(res)
}

func divideGeneric[T Number](x, y T) result.Result[T] {
	if y == 0 {
		return result.Err[T](DivideZeroError)
	}

	res := x / y

	return result.Ok(res)
}
