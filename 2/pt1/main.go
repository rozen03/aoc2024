package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed data.txt
var data string

type List struct {
	Values []int
}

type Generator struct {
	data  []List
	index int
}

func NewGenerator() *Generator {
	var lists []List
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		if line != "" {
			parts := strings.Fields(line)
			lista := make([]int, len(parts))
			for i, part := range parts {
				var a int
				_, err := fmt.Sscanf(part, "%d", &a)
				if err != nil {
					panic(err)
				}
				lista[i] = a
			}
			lists = append(lists, List{Values: lista})
		}
	}

	return &Generator{
		data:  lists,
		index: 0,
	}
}

func (g *Generator) Next() (*List, bool) {
	if g.index >= len(g.data) {
		return nil, false // No more pairs to return
	}

	pair := g.data[g.index]
	g.index++
	return &pair, true
}

func IsDecrecent(list []int) bool {
	for i := 1; i < len(list); i++ {
		if list[i] < list[i-1] {
			return false
		}
		if !HasDifferentOneTwoThree(list[i], list[i-1]) {
			return false
		}
	}
	return true
}

func IsIncrescent(list []int) bool {
	for i := 1; i < len(list); i++ {
		if list[i] > list[i-1] {
			return false
		}
		if !HasDifferentOneTwoThree(list[i], list[i-1]) {
			return false
		}
	}
	return true
}

func HasDifferentOneTwoThree(a, b int) bool {
	diff := a - b
	newDiff := math.Abs(float64(diff))
	if newDiff < 1 || newDiff > 3 {
		return false
	}
	return true
}
func main() {
	gen := NewGenerator()
	safeCount := 0
	for {
		list, ok := gen.Next()
		if !ok {
			break
		}
		if !IsDecrecent(list.Values) && !IsIncrescent(list.Values) {
			continue
		}

		safeCount++
	}

	fmt.Println("the count is:", safeCount)

}
