package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

func (e gasEngine) gasSearch() uint8 {
	return e.gallons * e.mpg
}

func canMakeit(e gasEngine, miles unit8) {
	if miles <= e.gasSearch() {
		fmt.Println("Can get the distance")
	} else {
		fmt.Println("Can`t get")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{25, 35}
	fmt.Println(myEngine.mpg, myEngine.gallons)
	fmt.Printf("The total gasSearch is %v\n", myEngine.gasSearch())
}
