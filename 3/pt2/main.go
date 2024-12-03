package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed data.txt
var data string

func main() {
	sum := 0

	dos := strings.Split(data, "do()")
	for _, do := range dos {
		donts := strings.Split(do, "don't()")
		fmt.Println("do()", donts[0])
		habilitaDO := donts[0]
		muls := strings.Split(habilitaDO, "mul(")
		for _, mul := range muls {
			var a, b int
			_, err := fmt.Sscanf(mul, "%d,%d)", &a, &b)
			if err != nil {
				continue
			}

			fmt.Println("mul(", mul, a*b)
			sum += a * b
		}

	}
	fmt.Println("sum", sum)
}