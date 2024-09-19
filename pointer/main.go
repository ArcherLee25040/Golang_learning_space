package main

import "fmt"

func main() {
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("The location is %p", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("%v\n", result)
	fmt.Printf("%v\n", thing1)
}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("%p\n", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]

	}
	return *thing2
}
