// Test for interface
// 通过接口实现多个方法复用
package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

type electronEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electronEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

//interface 定义milesLeft()方法可以被多个对象调用
type engine interface {
	milesLeft() uint8
}

func canMakeit(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can get to the destination")
	} else {
		fmt.Println("You can`t get to the destination,need to fuel up first")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{25, 35}
	canMakeit(myEngine, 50)
}
