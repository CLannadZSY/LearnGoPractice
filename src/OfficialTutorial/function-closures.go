/*

函数的闭包
Go 函数可以是一个闭包。闭包是一个函数值，它引用了其函数体之外的变量。
该函数可以访问并赋予其引用的变量的值，换句话说，该函数被这些变量“绑定”在一起。

例如，函数 adder 返回一个闭包。每个闭包都被绑定在其各自的 sum 变量上。
*/

package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func AdderChange(x int) int {
	sum := 0
	sum += x
	return sum
}

func main() {
	pos, neg := adder(), adder()
	PosAdd, NegAdd := AdderChange, AdderChange

	for i := 0; i < 10; i++ {
		fmt.Println("i=", i, pos(i), neg(2*i), PosAdd(i), NegAdd(2*i))
	}

}
