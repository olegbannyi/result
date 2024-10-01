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

	res := devideInt(6, 2).Expect("You won't see this message")

	fmt.Printf("Result: %v\n", res)

	resInt := devideGeneric(10, 2).Expect("You won't see this message")
	fmt.Printf("Result: %v\n", resInt)

	resFloat := devideGeneric(3.6, 2).Expect("You won't see this message")
	fmt.Printf("Result: %v\n", resFloat)

	resGen := devideGeneric(3.6, 0).Expect("Some custom message on devision by zero")
	fmt.Printf("Result: %v\n", resGen)
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
