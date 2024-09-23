// This package is for testing sumSlice 泛型
package main

import "fmt"

func main() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice[int](intSlice))

	var float32Slice = []float32{1, 2, 3}
	fmt.Println(sumSlice[float32](float32Slice))
}

func sumSlice[T any](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum

}
