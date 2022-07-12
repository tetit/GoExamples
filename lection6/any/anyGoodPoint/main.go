package main

import (
	"errors"
	"fmt"
	"strconv"
)

type StringVal string

func (sv StringVal) String() string {
	return string(sv)
}

type IntVal int

func (iv IntVal) String() string {
	return strconv.Itoa(int(iv))
}

func main() {
	inputs := []any{
		StringVal("hello"),
		IntVal(123),
		true,
		3.14,
	}

	var resVal string
	var err error

	for _, inputValue := range inputs {

		// resVal, err = AssertInterface(inputValue)
		// fmt.Printf("\nAsserion function: %s\n", "AssertInterface")
		// fmt.Printf("inputType: %[1]T | inputValue: %[1]v\n", inputValue)
		// fmt.Printf("outputType: %[1]T | outputValue: %[1]v\n", resVal)
		// fmt.Printf("error value: %v\n", err)

		resVal, err = AssertInterfaceVarying(inputValue)
		fmt.Printf("\nAsserion function: %s\n", "AssertStringVarying")
		fmt.Printf("inputType: %[1]T | inputValue: %[1]v\n", inputValue)
		fmt.Printf("outputType: %[1]T | outputValue: %[1]v\n", resVal)
		fmt.Printf("error value: %v\n", err)
	}

}

// func AssertInterface(v any) (string, error) {
// 	stringer, ok := v.(fmt.Stringer)
// 	if !ok {
// 		return "", errors.New("not supported type")
// 	}
// 	return stringer.String(), nil

// }

func AssertInterfaceVarying(v any) (string, error) {
	switch convertVal := v.(type) {
	case fmt.Stringer:
		return convertVal.String(), nil
	case bool:
		return strconv.FormatBool(convertVal), nil
	default:
		return "", errors.New("not supported type")

	}
}
