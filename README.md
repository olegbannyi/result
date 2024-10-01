# Result

Wrapping a function's return value to avoid direct handling of "nil" and to increase flexibility in error handling.
Inspired by [Rust error handling](https://doc.rust-lang.org/std/result/)

## Example of some math function - division
```go
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

type Number interface {
	int | float64
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
```
## Examples of usage of `result.Result[T]` for given math function:
### `result.Result[T].Unwrap()`
```go
// Success
res := divideInt(6, 2).Unwrap()
fmt.Printf("%v\n", res) // 3

// Error
res = divideInt(6, 0).Unwrap() // panic("cannot be divided by zero")
```
### `result.Result[T].UnwrapOr(T)`
```go
// Success
res := divideInt(6, 2).UnwrapOr(-1)
fmt.Printf("%v\n", res) // 3

// Error
res = divideInt(6, 0).UnwrapOr(-1)
fmt.Printf("%v\n", res) // -1
```
### `result.Result[T].UnwrapOrDefault()`
```go
// Success
res := divideGeneric(6, 2).UnwrapOrDefault()
fmt.Printf("%v\n", res) // 3

// Error
res = divideGeneric(6, 0).UnwrapOrDefault()
fmt.Printf("%v\n", res) // 0
```
### `result.Result[T].UnwrapOrElse(f func() T)`
```go
// Success
res := divideGeneric(6, 2).UnwrapOrElse(func(){ return -1 })
fmt.Printf("%v\n", res) // 3

// Error
res = divideGeneric(6, 0).UnwrapOrElse(func(){ return -1 })
fmt.Printf("%v\n", res) // -1
```
### `result.Result[T].Expect(string)`
```go
// Success
res := divideInt(6, 2).Expect("invalid values")
fmt.Printf("%v\n", res) // 3

// Error
res = divideInt(6, 0).Expect("invalid values") // panic("invalid values")
```
### `result.Result[T].IsOk()`
```go
// Success
res := divideInt(6, 2).IsOk()
fmt.Printf("%v\n", res) // true

// Error
res = divideInt(6, 0).IsOk() 
fmt.Printf("%v\n", res) // false
```
### `result.Result[T].IsErr()`
```go
// Success
res := divideGeneric(6, 2).IsErr()
fmt.Printf("%v\n", res) // false

// Error
res = divideGeneric(6, 0).IsErr() 
fmt.Printf("%v\n", res) // true
```