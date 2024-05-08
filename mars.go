package main

import (
	"fmt"
	"math/rand"

	"github.com/charmbracelet/bubbles/table"
)

var spaceLines []string = []string{"Virgin Galactic", "SpaceX", "Space Adventures"}

const distanceToEarth = 62100000 // km
const secondsPerDay = 60 * 60 * 24

const minSpeed = 16 // km/s
const maxSpeed = 30 // km/s

const minPrice = 36 // MM
const maxPrice = 50 // MM

type SpaceTrip struct {
	price     int
	speed     int
	roundTrip bool
	spaceline string
}

func RandomPrice() int {
	return rand.Intn(maxPrice-minPrice) + minPrice
}

func RandomSpeed() int {
	return rand.Intn(maxSpeed-minSpeed) + minSpeed
}

func RandomRoundTrip() bool {
	return rand.Intn(2) == 1
}

func RandomSpaceLine() string {
	var max = len(spaceLines)
	i := rand.Intn(max - 1)
	return spaceLines[i]
}

func (s SpaceTrip) Price() int {
	if s.roundTrip {
		return s.price * 2
	} else {
		return s.price
	}
}

func (s SpaceTrip) GetDuractionDays() int {
	return distanceToEarth / s.speed / secondsPerDay
}

func NewSpaceTrip() SpaceTrip {
	return SpaceTrip{
		price:     RandomPrice(),
		speed:     RandomSpeed(),
		roundTrip: RandomRoundTrip(),
		spaceline: RandomSpaceLine(),
	}
}

func (s SpaceTrip) ToRow() table.Row {
	var tripType string

	if s.roundTrip {
		tripType = "Round-trip"
	} else {
		tripType = "One-way"
	}

	// spaceline, days, trip type, price
	r := table.Row{s.spaceline, fmt.Sprint(s.GetDuractionDays()), tripType, "$" + fmt.Sprint(s.Price())}
	return r
}
