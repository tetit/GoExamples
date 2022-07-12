package main

import (
	"errors"
	"fmt"
)

type transport interface {
	avgSpeedGetter
	fmt.Stringer
}

type avgSpeedGetter interface {
	getAvgSpeed() float64
}

func minutesInRoad(asg avgSpeedGetter, pointA, pointB int) (float64, error) {
	if pointA > pointB {
		return 0, errors.New("wrong input data")
	}
	avgSpeed := asg.getAvgSpeed()

	if err := validateSpeed(avgSpeed); err != nil {

		return 0, fmt.Errorf("filed to calculate minutesInRoad: %w", err)
	}

	distance := pointB - pointA
	kmPerMinute := avgSpeed / 60
	roadTime := float64(distance) / kmPerMinute

	return roadTime, nil
}

var slowTransportErr = errors.New("it is to slow")

func validateSpeed(avgSpeed float64) error {

	if avgSpeed < 80 {
		return slowTransportErr
		// return fmt.Errorf("it is to slow %s", "oups")
		// return errors.New("It is to slow")
	}
	return nil
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
		minT, err := minutesInRoad(t, pointA, pointB)
		if errors.Is(err, slowTransportErr) {
			fmt.Printf("\n%s in not option: %v", t, err)
			// fmt.Printf("\n%s in not option %s", t, err.Error())
			continue
		}
		if err != nil {
			fmt.Printf("all is bad: %v", err)
			return

		}
		fmt.Printf("\nMinutes in road by %s: %f", t /*.String()*/, minT)
	}

}
