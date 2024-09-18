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
	//键对
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap_2 = map[string]uint8{"Admin": 23, "Search": 45}
	fmt.Println(myMap_2["Search"])
	var age, ok = myMap_2["Admin"]
	delete(myMap_2, "Admin")
	if ok {
		fmt.Println("The age is %v", age)
	} else {
		fmt.Println("Invalid name")
	}

	for name, age := range myMap_2 {
		fmt.Printf("Name: %v,age: %v\n", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v,Value:%v\n", i, v)
	}
	//git
}
