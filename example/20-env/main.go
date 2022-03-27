package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// go run example/20-env/main.go a b c d
	fmt.Println(os.Args)           // [/var/folders/8p/n34xxfnx38dg8bv_x8l62t_m0000gn/T/go-build3406981276/b001/exe/main a b c d]
	fmt.Println(os.Getenv("PATH")) // /usr/local/go/bin...
	fmt.Println(os.Setenv("AA", "BB"))

	buf, err := exec.Command("grep", "127.0.0.1", "/etc/hosts").CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf)) // 127.0.0.1       localhost
}
