package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func add2(a, b int) int { // 参数a没有指定类型
	return a + b // b类型为int，根据a+b推断出a也为int类型
}

func exists(m map[string]string, k string) (v string, ok bool) {
	v, ok = m[k]
	return v, ok
}

func main() {
	res := add(1, 2)
	fmt.Println(res) // 3

	res2 := add2(1, 2)
	fmt.Println(res2) // 3

	v, ok := exists(map[string]string{"a": "A"}, "a")
	fmt.Println(v, ok) // A True
}
