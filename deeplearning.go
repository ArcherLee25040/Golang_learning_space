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
	fmt.Println(len("test"))

	var myBoolean bool = true
	fmt.Println(myBoolean)

	myString1 := "hello"
	fmt.Println(myString1)

	var int1, int2 = 1, 2
	fmt.Println(int1, int2)
}
