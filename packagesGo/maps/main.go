package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

// var m1 = map[string]Vertex{
// 	"Bell Labs": Vertex{
// 		40.68433, -74.39967,
// 	},
// 	"Google": Vertex{
// 		37.42202, -122.08408,	
// 	},
// }

var m1 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m)
	fmt.Println(m1)

	v, ok := m1["Google"]
	fmt.Println("The value:", v, "Present?", ok)

	
}