package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	n := 0
	input := bufio.NewScanner(os.Stdin)

	//调用一次读取一行
	input.Scan()
	str1 := input.Text()
	str1 = strings.ToLower(str1)

	//读取第二行
	input.Scan()
	str2 := input.Text()
	str2 = strings.ToLower(str2)

	for _, value := range str1 {
		if str2 == string(value) {
			n++
		}
	}
	fmt.Println(n)

}
