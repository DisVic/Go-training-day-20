package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b int
	_, _ = fmt.Scan(&a, &b)
	checkNumbers(a, b)
}

func checkNumbers(a, b int) {
	str1 := strconv.Itoa(a)
	str2 := strconv.Itoa(b)
	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			if str1[i] == str2[j] {
				fmt.Print(string(str1[i]), " ")
				break
			}
		}
	}
}
