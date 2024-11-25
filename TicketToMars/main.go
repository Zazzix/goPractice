package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var companies = []string{"SpaceX", "Space Adventures", "Virgin Galactic"}
	var tripType = []string{"One-way", "Round-trip"}
	const secondsPerDay = 86400
	minSpeed := 16
	maxSpeed := 30
	distance := 62100000

	for i := 0; i < 10; i++ {
		carrier := companies[rand.Intn(len(companies))]
		trip := tripType[rand.Intn(len(tripType))]
		speed := (rand.Intn(maxSpeed-minSpeed+1) + minSpeed)
		tripDuration := (distance / (speed * 3600)) / 24
		tripPrice := 20 + speed

		if trip == "Round-trip" {
			tripPrice *= 2
		}

		fmt.Printf("%-16v %4v %-10v $%4v\n", carrier, tripDuration, trip, tripPrice)
	}
}
