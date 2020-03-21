package chart

import (
	"fmt"
	"math/rand"
)

func ExamplePolyline() {
	rand.Seed(0)
	data := make([]int, 100)
	for i := range data {
		data[i] = rand.Intn(200) + 50
	}
	// ioutil.WriteFile("./_doc/demo.svg", []byte(Polyline(data)), 0644)
	// ioutil.WriteFile("./_doc/demo.svg", []byte(Polyline([]int{5, 7, 14, 23, 45, 5, 40})), 0644)
	fmt.Println("ok")

	// Output:
	// ok
}
