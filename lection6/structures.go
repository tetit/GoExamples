package main

type truck struct {
	capacity int
	avgSpeed float64
	height   int
}

func (t truck) getAvgSpeed() float64 {
	return t.avgSpeed
}

func (t truck) getTransportName() string {
	return "truck"
}


type train struct {
	capacity int
	avgSpeed float64
	carts    int
}

func (t train) getAvgSpeed() float64 {
	return t.avgSpeed
}

func (t train) getTransportName() string {
	return "train"
}

type pickup struct {
	capacity int
	avgSpeed float64
	gasTank  int
}

func (p pickup) getAvgSpeed() float64 {
	return p.avgSpeed
}

func (p pickup) getTransportName() string {
	return "pickup"
}
