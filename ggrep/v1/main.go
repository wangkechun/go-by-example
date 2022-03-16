package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func find(pattern string, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	fileReader := bufio.NewReader(file)
	lineIdx := 0
	for {
		line, err := fileReader.ReadString('\n')
		if err == io.EOF {
			break
		}
		lineIdx++
		if strings.Contains(line, pattern) {
			fmt.Printf("%v %v:\t%v\n", filename, lineIdx, strings.Trim(line, "\n"))
		}
	}
	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("ggrep pattern filename")
		os.Exit(1)
	}
	err := find(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Printf("find error: %+v\n", err)
	}
}
