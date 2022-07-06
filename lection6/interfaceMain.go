package main

import "fmt"

type transport interface {
	avgSpeedGetter
	transportNameGetter
}

type avgSpeedGetter interface {
	getAvgSpeed() float64
}

type transportNameGetter interface {
	getTransportName() string
}

func minutesInRoad(asg avgSpeedGetter, pointA, pointB int) float64 {
	distance := pointB - pointA
	kmPerMinute := asg.getAvgSpeed() / 60
	roadTime := float64(distance) / kmPerMinute
	return roadTime
}

func main() {

	pointA, pointB := 5, 30

	// transports := map[string]avgSpeedGetter{
	// 	"truck":  truck{avgSpeed: 70},
	// 	"train":  train{avgSpeed: 100},
	// 	"pickup": pickup{avgSpeed: 80},
	// }

	transports := []transport{
		truck{avgSpeed: 70},
		train{avgSpeed: 100},
		pickup{avgSpeed: 80},
	}

	for _, t := range transports {
		minT := minutesInRoad(t, pointA, pointB)
		fmt.Printf("\nMinutes in road by %s: %f", t.getTransportName(), minT)
	}

}
