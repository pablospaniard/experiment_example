package main

import (
	"errors"
	"math/rand"
)

// generator generates ints and save them in local memory.
type generator struct {
	rand  *rand.Rand
	store [][]int
}

func newGenerator(seed int64) *generator {
	return &generator{
		rand:  rand.New(rand.NewSource(seed)),
		store: make([][]int, 0),
	}
}

// save generates integers and save them in memory, returning their index.
func (g *generator) save(size int) (int, error) {
	nums := g.generate(size)
	g.store = append(g.store, nums)

	return len(g.store) - 1, nil
}

func (g *generator) generate(size int) []int {
	numbers := make([]int, size)

	for i := range numbers {
		numbers[i] = g.rand.Int()
	}

	return numbers
}

func (g *generator) fetch(index int) ([]int, error) {
	if index < 0 || index > len(g.store)-1 {
		return nil, errors.New("index not valid")
	}

	return g.store[index], nil
}
