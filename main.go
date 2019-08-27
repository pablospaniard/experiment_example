package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/jelmersnoeck/experiment"
)

func main() {
	oldGen, newGen := newGenerator(2), newGenerator(4)

	exp := experiment.New(
		experiment.WithPercentage(50),
		// All your "metadata" regarding the experiment will go here as settings for the publisher.
		experiment.WithPublisher(experiment.NewLogPublisher(fmt.Sprintf("generation-%s", time.Now()), nil)),
	)

	exp.Control(func() (interface{}, error) {
		index, err := oldGen.save(3)
		if err != nil {
			return nil, err
		}

		fetch, err := oldGen.fetch(index)

		return fetch, err
	})

	exp.Candidate("candidate1", func() (interface{}, error) {
		index, err := newGen.save(3)
		if err != nil {
			return nil, err
		}

		fetch, err := newGen.fetch(index)

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
