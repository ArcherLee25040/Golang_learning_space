package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "hello"
	printMe(printValue)

	var numbertor int = 11
	var denominator int = 22
	var result, remainder, err = intDivision(numbertor, denominator)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Println("The result of the integer division is %v with remainder %v", result, remainder)
	// } else {
	// 	fmt.Println("The result is %v with remainer %v", result, remainder)
	// }
	switch {
	case err != nil:
		fmt.Print(err.Error())
	case remainder == 0:
		fmt.Print("The result of the integer division is %v", result)
	default:
		fmt.Print("The result is %v with remainder %v", result, remainder)
	}
	fmt.Println(result)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numbertor int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannotdivide by zero")
		return 0, 0, err
	}
	var result int = numbertor / denominator
	var remainder int = numbertor % denominator
	return result, remainder, err
}
