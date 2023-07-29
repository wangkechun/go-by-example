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
}
