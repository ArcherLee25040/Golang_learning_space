package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

func (e gasEngine) gasSearch() uint8 {
	return e.gallons * e.mpg
}

func main() {
	var myEngine gasEngine = gasEngine{25, 35}
	fmt.Println(myEngine.mpg, myEngine.gallons)
	fmt.Printf("The total gasSearch is %v\n", myEngine.gasSearch())
}
