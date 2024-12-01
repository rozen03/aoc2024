package main

import (
	_ "embed"
	"fmt"
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
	bCounter := 0
	prevResult := 0.0
	prevValue := 0
	for _, a := range listA {
		//use reused calc if the value is the same as this is a sorted list
		if a == prevValue {
			counter += prevResult
			continue
		}
		// ignore values that don't matter, ihey are less than a skip them
		for listB[bCounter] < a {
			bCounter++
		}
		//count how many appearances are of the a value
		prevCounter := bCounter
		for a == listB[bCounter] {
			bCounter++
		}
		//result
		appearances := bCounter - prevCounter
		prevResult = float64(a * appearances)
		prevValue = a

		counter += prevResult

	}
	// Your code here
	return int(counter)
}
