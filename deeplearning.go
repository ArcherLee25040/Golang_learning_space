package main

import "fmt"

func main() {
	var intNum int32 = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float32 = 12321
	fmt.Println(floatNum)

	var myString string = "Hello \nworld!"
	fmt.Println(myString)
}
