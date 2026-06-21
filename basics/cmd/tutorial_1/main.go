package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func main() {

	/*
		SIGNED INTEGERS
		int     - -2147483648 to 2147483647 (32 bit system, 2^32 total values) / -9223372036854775808 to 9223372036854775807 (64 bit system, 2^64 total values)
		int8    - -128 to 127 (256 total values)
		int16   - -32768 to 32767 (2^16 total values)
		int32   - -2147483648 to 2147483647 (2^32 total values)
		int64   - -9223372036854775808 to 9223372036854775807 (2^64 total values)

		UNSIGNED INTEGERS
		uint	- 0 to 4294967295 (32 bit system) / 0 to 18446744073709551615 (64 bit system)
		uint8	- 0 to 255
		uint16	- 0 to 65535
		uint32	- 0 to 4294967295
		uint64	- 0 to 18446744073709551615

		default value is 0 for all
	*/
	var intNum int = 100
	fmt.Println(intNum)

	/*
		float32				- (32 bits) - -3.4e+38 to 3.4e+38.
		float64 (default)	- (64 bits) - -1.7e+308 to +1.7e+308.

		default value is 0 for both
	*/
	var floatNum float64 = 100.56
	fmt.Println(floatNum)

	/*
		STRING
		default value for string is ''
	*/
	fmt.Println("Hello World!")

	// String length
	fmt.Println(len("A@")) // this returns the number of bytes. not the number of characters.

	fmt.Println(utf8.RuneCountInString("A@")) // rune: data type for characters

	/*
		BOOLEAN

		default value is false
	*/
	var myBoolean bool = false
	fmt.Println(myBoolean)

	const myConst string = "const value" // cannot be changed later and should always have a value assigned

	// =================================================

	// FUNCTIONS & CONSTROL STRUCTURES
	var printValue string = "Hello world"
	printMe(printValue)

	var numerator int = 1
	var denominator int = 1
	var result, remainder, err = intDivision(numerator, denominator)

	// handle using if else statements
	fmt.Println("Handle with if else")
	if err != nil {
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v \n", result)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v \n", result, remainder)
	}

	// handle using switch
	fmt.Println("Handle with switch")
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v \n", result)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v \n", result, remainder)
	}

	// handle using switch with a variable
	fmt.Println("Handle with switch with a variable")
	switch remainder {
	case 0:
		fmt.Printf("The result of the integer division is %v \n and the remainder is 0 \n", result)
	case 1, 2:
		fmt.Printf("The result of the integer division is %v \n and the remainder is 1 or 2 \n", result)
	default:
		fmt.Printf("The result of the integer division is %v \n and the remainder is %v \n", result, remainder)
	}

	fmt.Printf("The result of the integer divisin is %v with remainder %v \n", result, remainder)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot Devide by Zero")
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
