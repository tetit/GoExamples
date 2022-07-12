package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	inputs := []any{
		"hello",
		123,
		true,
		3.14,
	}

	var resVal string
	var err error

	for _, inputValue := range inputs {
		// resVal = AssertString(inputValue)
		// fmt.Printf("\nAsserion function: %s\n", "AssertString")
		// fmt.Printf("inputType: %[1]T | inputValue: %[1]v\n", inputValue)
		// fmt.Printf("outputType: %[1]T | outputValue: %[1]v\n", resVal)

		// resVal, err = AssertStringOkNotation(inputValue)
		// fmt.Printf("\nAsserion function: %s\n", "AssertStringOkNotation")
		// fmt.Printf("inputType: %[1]T | inputValue: %[1]v\n", inputValue)
		// fmt.Printf("outputType: %[1]T | outputValue: %[1]v\n", resVal)
		// fmt.Printf("error value: %v\n", err)

		resVal, err = AssertStringVarying(inputValue)
		fmt.Printf("\nAsserion function: %s\n", "AssertStringVarying")
		fmt.Printf("inputType: %[1]T | inputValue: %[1]v\n", inputValue)
		fmt.Printf("outputType: %[1]T | outputValue: %[1]v\n", resVal)
		fmt.Printf("error value: %v\n", err)
	}

	fmt.Println(string(123))
}

// func AssertString(v any/*interface{}*/) string {  //not goog code
// 	stringVal := v.(string)
// 	return stringVal

// }

// func AssertStringOkNotation(v interface{}) (string, error) {
// 	stringVal, ok := v.(string)
// 	if !ok {
// 		return "", errors.New("not supported type")
// 	}
// 	return stringVal, nil

// }

func AssertStringVarying(v interface{}) (string, error) {
	switch convertVal := v.(type) {
	case string:
		return convertVal, nil
	case int:
		return strconv.Itoa(convertVal), nil
	case bool:
		return strconv.FormatBool(convertVal), nil
	default:
		return "", errors.New("not supported type")

	}

}
