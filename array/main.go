// this file is for array_test
package main

import "fmt"

func main() {
	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
	var intslice []int32 = []int32{1, 2, 3}
	var intslice_2 []int32 = []int32{4, 5, 6}
	intslice = append(intslice, intslice_2...)
	fmt.Println(intslice)

}
