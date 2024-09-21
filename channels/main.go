// 定义功能：定义网站监听商品属性变化情况
package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

var MAX_Chicken_price float32 = 5
var MAX_TOFU_price float32 = 3

func main() {
	var chickenchannel = make(chan string)
	var tofuchannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites {
		go Check_Chicken_Prices(websites[i], chickenchannel)
		go Check_Tofu_Prices(websites[i], tofuchannel)
	}
	sendMessage(chickenchannel, tofuchannel)
}

func Check_Tofu_Prices(websites string, tofuchannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofu_price = rand.Float32() * 20
		if tofu_price <= MAX_TOFU_price {
			tofuchannel <- websites
			break
		}
	}

}

func Check_Chicken_Prices(websites string, chickenchannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chicken_price = rand.Float32() * 20
		if chicken_price <= MAX_Chicken_price {
			chickenchannel <- websites
			break
		}
	}

}

func sendMessage(chickenchannel chan string, tofuchannel chan string) {
	// fmt.Printf("\nFound a deal on chicken at %s", <-chickenchannel)
	// fmt.Printf("\nFound a deal on chicken at %s", <-tofuchannel)
	select {
	case websites := <-chickenchannel:
		fmt.Printf("Found deal on chicken at %v\n", websites)
	case websites := <-tofuchannel:
		fmt.Printf("Found deal on chicken at %v\n", websites)
	}
}
