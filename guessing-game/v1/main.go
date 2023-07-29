package main

import (
	"fmt"
	"math/rand"
)

func main() {
	maxNum := 100
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)
	// 新版本（我的为1.20.6）已经可以每次输出不同的数了
	// 查阅资料得，从1.10版本开始
	// math/rand 包会自动初始化一个全局公共源，并默认使用当前时间作为种子，因此不再需要手动设置种子值。
}
