package main

import "fmt"

//此时你可以去掉分号，因为 C 的 while 在 Go 中叫做 for。
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	main2()
}

//无限循环
//如果省略循环条件，该循环就不会结束，因此无限循环可以写得很紧凑。
func main2() {
	a := 1
	for {
		fmt.Println(a)
	}
}
