package main

import "fmt"

func main() {
	var c = make(chan int)
	go process(c)
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
}

func process(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
	}

}
