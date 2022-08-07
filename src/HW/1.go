package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	input.Scan()
	a := input.Text()

	array := strings.Fields(a)

	fmt.Println(len(array[len(array)-1]))

}
