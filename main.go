package main

import (
	"fmt"
	"reflect"

	"github.com/jelmersnoeck/experiment"
)

func main() {
	g1, g2 := newGenerator(2), newGenerator(4)

	exp := experiment.New(
		experiment.WithPercentage(50),
		experiment.WithPublisher(experiment.NewLogPublisher("generator", nil)),
	)

	exp.Control(func() (interface{}, error) {
		index, err := g1.save(3)
		fetch, _ := g1.fetch(index)

		return fetch, err
	})

	exp.Candidate("candidate1", func() (interface{}, error) {
		index, err := g2.save(3)
		fetch, _ := g2.fetch(index)

		return fetch, err
	})

	exp.Force(true)

	exp.Compare(func(a, b interface{}) bool {
		g1Results := a.([]int)
		g2Results := b.([]int)
		return reflect.DeepEqual(g1Results, g2Results)
	})

	result, err := exp.Run()
	if err != nil {
		panic(err)
	} else {
		fmt.Println(result)
	}
}
