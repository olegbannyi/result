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

	res := divideInt(6, 2).Expect("You won't see this message")

	fmt.Printf("Result: %v\n", res)

	resInt := divideGeneric(10, 2).Expect("You won't see this message")
	fmt.Printf("Result: %v\n", resInt)

	resFloat := divideGeneric(3.6, 2).Expect("You won't see this message")
	fmt.Printf("Result: %v\n", resFloat)

	resGen := divideGeneric(3.6, 0).Expect("Some custom message on devision by zero")
	fmt.Printf("Result: %v\n", resGen)
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
