package result_test

import (
	"errors"
	"testing"

	"gotest.tools/assert"

	"github.com/olegbannyi/result"
)

func TestOk(t *testing.T) {
	foo := struct {
		val string
	}{"bar"}

	tests := []struct {
		actual any
		want   any
	}{
		{
			actual: result.Ok(1),
			want:   result.NewResult(1, nil),
		},
		{
			actual: result.Ok(true),
			want:   result.NewResult(true, nil),
		},
		{
			actual: result.Ok(foo),
			want:   result.NewResult(foo, nil),
		},
	}

	for _, test := range tests {
		if test.actual != test.want {
			t.Errorf("actual = %#v, want = %#v", test.actual, test.want)
		}
	}
}

func TestErr(t *testing.T) {
	type foo struct{}

	err := errors.New("some error")

	tests := []struct {
		actual any
		want   any
	}{
		{
			actual: result.Err[int](err),
			want:   result.NewResult(0, err),
		},
		{
			actual: result.Err[bool](err),
			want:   result.NewResult(false, err),
		},
		{
			actual: result.Err[foo](err),
			want:   result.NewResult(foo{}, err),
		},
	}

	for _, test := range tests {
		if test.actual != test.want {
			t.Errorf("actual = %#v, want = %#v", test.actual, test.want)
		}
	}
}

func TestAsTuple(t *testing.T) {
	someVal := struct{}{}
	someErr := errors.New("some error")

	gotVal, gotErr := result.NewResult(someVal, someErr).AsTuple()
	if gotVal != someVal {
		t.Errorf("actual value => %v, want value => %v", gotVal, someVal)
	}
	if gotErr != someErr {
		t.Errorf("actual error => %v, want error => %v", gotErr, someErr)
	}
}

func TestUnwrapOr(t *testing.T) {
	someErr := errors.New("some error")

	// Example with integers
	someInt := 5
	someDefaultInt := 7

	gotInt := result.NewResult(someInt, nil).UnwrapOr(someDefaultInt)
	assert.Equal(t, someInt, gotInt)

	gotInt = result.NewResult(0, someErr).UnwrapOr(someDefaultInt)
	assert.Equal(t, someDefaultInt, gotInt)

	// Example with strings
	someString := "bar"
	someDefaultStrint := "foo"

	gotString := result.NewResult(someString, nil).UnwrapOr(someDefaultStrint)
	assert.Equal(t, someString, gotString)

	gotString = result.NewResult("", someErr).UnwrapOr(someDefaultStrint)
	assert.Equal(t, someDefaultStrint, gotString)

	// Examlpe with bool
	someBool := false
	someDefaultBool := true

	gotBool := result.NewResult(someBool, nil).UnwrapOr(someDefaultBool)
	assert.Equal(t, someBool, gotBool)

	gotBool = result.NewResult(false, someErr).UnwrapOr(someDefaultBool)
	assert.Equal(t, someDefaultBool, gotBool)

	// Example with struct
	someStruct := struct{ val string }{"bar"}
	someDefaultStruct := struct{ val string }{"foo"}

	gotStruct := result.NewResult(someStruct, nil).UnwrapOr(someDefaultStruct)
	assert.Equal(t, someStruct, gotStruct)

	gotStruct = result.NewResult(struct{ val string }{}, someErr).UnwrapOr(someDefaultStruct)
	assert.Equal(t, someDefaultStruct, gotStruct)
}

func TestIsOk(t *testing.T) {
	someErr := errors.New("some error")

	isOk := result.NewResult(struct{}{}, nil).IsOk()
	assert.Assert(t, isOk)

	isOk = result.NewResult(struct{}{}, someErr).IsOk()
	assert.Assert(t, !isOk)
}

func TestIsErr(t *testing.T) {
	someErr := errors.New("some error")

	isErr := result.NewResult(struct{}{}, nil).IsErr()
	assert.Assert(t, !isErr)

	isErr = result.NewResult(struct{}{}, someErr).IsErr()
	assert.Assert(t, isErr)
}

func TestUnwrapSuccess(t *testing.T) {
	someVal := struct{ name string }{"foo"}

	gotVal := result.NewResult(someVal, nil).Unwrap()
	assert.Equal(t, gotVal, someVal)
}

func TestUnwrapFailure(t *testing.T) {
	errorMsg := "some error"
	someErr := errors.New(errorMsg)

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Error got: %#v, want: %s", err, someErr.Error())
		} else {
			assert.Equal(t, someErr, err)
		}
	}()

	_ = result.NewResult(struct{}{}, someErr).Unwrap()
}

func TestExpectSuccess(t *testing.T) {
	someVal := struct{ name string }{"foo"}
	someErrorMsg := "some error"

	gotVal := result.NewResult(someVal, nil).Expect(someErrorMsg)
	assert.Equal(t, gotVal, someVal)
}

func TestExpectFailure(t *testing.T) {
	someCustomErrorMsg := "some custom error message"
	someErr := errors.New("some error")

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("Error got: %#v, want: %s", err, someCustomErrorMsg)
		} else {
			assert.Equal(t, someCustomErrorMsg, err)
		}
	}()

	_ = result.NewResult(struct{}{}, someErr).Expect(someCustomErrorMsg)
}
