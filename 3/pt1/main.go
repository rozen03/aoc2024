package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed data.txt
var data string

func main() {
	muls := strings.Split(data, "do")
	sum := 0
	for _, mul := range muls {
		var a, b int
		_, err := fmt.Sscanf(mul, "%d,%d)", &a, &b)
		if err != nil {
			continue
		}

		fmt.Println("mul(", mul, a*b)
		sum += a * b
	}
	fmt.Println("sum", sum)

}
