package main

import (
	"fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f) // 1.234

	n, _ := strconv.ParseInt("111", 10, 64)
	fmt.Println(n) // 111

	n, _ = strconv.ParseInt("0x1000", 0, 64) // 基数为零表示自动识别进制，这里通过0x前缀来判断为十六进制
	fmt.Println(n) // 4096

	n2, _ := strconv.Atoi("123") // Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.
	fmt.Println(n2) // 123

	n2, err := strconv.Atoi("AAA")
	fmt.Println(n2, err) // 0 strconv.Atoi: parsing "AAA": invalid syntax
}
