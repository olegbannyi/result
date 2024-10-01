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
	DevideZeroError MathOpError = 1 << iota
	SomeOtherError
)

func (e MathOpError) Error() string {
	switch e {
	case DevideZeroError:
		return "cannot be deived by zero"
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

	isOk := devideInt(6, 2).IsOk()
	isErr := devideInt(6, 2).IsErr()

	fmt.Printf("devideInt(6, 2) => isOk: %v, isErr: %v\n", isOk, isErr)

	isOk = devideInt(6, 0).IsOk()
	isErr = devideInt(6, 0).IsErr()

	fmt.Printf("devideInt(6, 0) => isOk: %v, isErr: %v\n", isOk, isErr)

	isOk = devideGeneric(10, 2).IsOk()
	isErr = devideGeneric(6, 2).IsErr()

	fmt.Printf("devideGeneric(10, 2) => isOk: %v, isErr: %v\n", isOk, isErr)

	isOk = devideGeneric(3.6, 2).IsOk()
	isErr = devideGeneric(3.6, 2).IsErr()

	fmt.Printf("devideGeneric(3.6, 2) => isOk: %v, isErr: %v\n", isOk, isErr)

	isOk = devideGeneric(3.6, 0).IsOk()
	isErr = devideGeneric(3.6, 0).IsErr()

	fmt.Printf("devideGeneric(3.6, 0) => isOk: %v, isErr: %v\n", isOk, isErr)
}

func devideInt(x, y int) result.Result[int] {
	if y == 0 {
		return result.Err[int](DevideZeroError)
	}

	res := x / y
	return result.Ok(res)
}

func devideGeneric[T Number](x, y T) result.Result[T] {
	if y == 0 {
		return result.Err[T](DevideZeroError)
	}

	res := x / y

	return result.Ok(res)
}
