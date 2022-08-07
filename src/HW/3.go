package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sss := bufio.NewScanner(os.Stdin)
	for sss.Scan() {
		str1 := sss.Text()
		count := len(str1)
		for i := 0; i < count; {
			if i+8 < count {
				fmt.Println(str1[i : i+8])

				i = i + 8

			} else {
				str2 := str1[i:]
				zeroStr := i + 8 - count
				for j := 0; j < zeroStr; j++ {
					str2=str2 + "0"
				}
				fmt.Println(str2)
				break

			}
		}

	}

}
