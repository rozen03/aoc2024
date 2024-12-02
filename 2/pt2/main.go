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
func copyAndDeleteIndex(list []int, index int) []int {
	newList := make([]int, 0, len(list)-1)
	for i := 0; i < len(list); i++ {
		if i == index {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

func main() {
	gen := NewGenerator()
	safeCount := 0
	for {
		list, ok := gen.Next()
		if !ok {
			break
		}

		if IsDecrecent(list.Values) || IsIncrescent(list.Values) {
			safeCount++
			continue
		}

		for i, _ := range list.Values {
			copia := copyAndDeleteIndex(list.Values, i)
			if IsDecrecent(copia) || IsIncrescent(copia) {
				safeCount++
				break
			}

		}
	}

	fmt.Println("the count is:", safeCount)

}
