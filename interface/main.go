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

/*
以下是对这段代码的详细讲解：

**一、整体结构和包声明**

1. `package main`声明这是一个可执行的主包。

**二、结构体定义**

1. `gasEngine`结构体：
   - 代表汽油发动机，包含两个字段`mpg`（每加仑英里数，表示燃油效率）和`gallons`（油箱中的加仑数），数据类型都是无符号 8 位整数`uint8`。
2. `electronEngine`结构体：
   - 代表电动发动机，包含两个字段`mpkwh`（每千瓦时英里数）和`kwh`（电池容量千瓦时数），数据类型也是`uint8`。

**三、方法定义**

1. 为`gasEngine`结构体定义了方法`milesLeft()`：
   - 该方法计算并返回汽油发动机剩余的行驶里程数，即油箱中的加仑数乘以每加仑英里数。
2. 为`electronEngine`结构体也定义了同名方法`milesLeft()`：
   - 计算并返回电动发动机剩余的行驶里程数，即电池容量千瓦时数乘以每千瓦时英里数。

**四、接口定义**

1. `engine`接口：
   - 只定义了一个方法`milesLeft()`，这意味着任何实现了这个方法的类型都可以被视为实现了`engine`接口。

**五、函数定义**

1. `canMakeit`函数：
   - 接收一个实现了`engine`接口的参数`e`和一个表示要行驶的英里数的参数`miles`。
   - 函数内部判断如果要行驶的英里数小于等于发动机剩余的行驶里程数（通过调用`e.milesLeft()`方法获得），则打印“你可以到达目的地”，否则打印“你不能到达目的地，需要先加油”。

**六、主函数**

1. `main`函数是程序的入口点。
   - 创建了一个`gasEngine`类型的变量`myEngine`，并初始化其`mpg`为 25，`gallons`为 35。
   - 调用`canMakeit`函数，传入`myEngine`和要行驶的英里数 50。根据代码逻辑，这里会判断剩余里程是否足够到达给定的行驶英里数，并打印相应的结果。

这段代码通过接口实现了多态性，使得不同类型的发动机可以通过实现同一个接口的方法来进行统一的操作（如判断是否能够到达目的地）。这样可以提高代码的灵活性和可扩展性，方便在后续添加更多类型的发动机而无需修改太多的现有代码。
*/
