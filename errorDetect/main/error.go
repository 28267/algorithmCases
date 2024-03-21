package main

import (
	"fmt"
	"time"
)

func Test(num1 int, num2 int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("fail err= ", err)
		}
	}()
	res := num1 / num2
	fmt.Println("结果是: ", res)
}

func main() {

	n1 := 10
	n2 := 0
	Test(n1, n2)
	time.Sleep(time.Second * 5)
	for {
		fmt.Println("代码继续向下执行..")
	}
}
