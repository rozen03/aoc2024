package main

import (
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strings"
)

//go:embed data.txt
var data string

type Pair struct {
	A int
	B int
}

type Generator struct {
	data  []Pair
	index int
}

func NewGenerator() *Generator {
	var pairs []Pair
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		if line != "" {
			parts := strings.Fields(line)
			if len(parts) == 2 {
				var a, b int
				_, err := fmt.Sscanf(parts[0], "%d", &a)
				if err != nil {
					fmt.Println("Error parsing integer:", err)
					continue
				}
				_, err = fmt.Sscanf(parts[1], "%d", &b)
				if err != nil {
					fmt.Println("Error parsing integer:", err)
					continue
				}
				pairs = append(pairs, Pair{A: a, B: b})
			}
		}
	}

	return &Generator{
		data:  pairs,
		index: 0,
	}
}

func (g *Generator) Next() (*Pair, bool) {
	if g.index >= len(g.data) {
		return nil, false // No more pairs to return
	}

	pair := g.data[g.index]
	g.index++
	return &pair, true
}

func main() {
	gen := NewGenerator()
	var listA []int
	var listB []int
	for {
		pair, ok := gen.Next()
		if !ok {
			break
		}
		listA = append(listA, pair.A)
		listB = append(listB, pair.B)

	}
	fmt.Println(solve(listA, listB))
}

func solve(listA, listB []int) int {
	slices.Sort(listA)
	slices.Sort(listB)
	counter := 0.0
	for i, a := range listA {
		b := listB[i]
		counter += math.Abs(float64(a - b))
	}
	// Your code here
	return int(counter)
}
