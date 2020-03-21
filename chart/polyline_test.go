package chart

import (
	"fmt"
	"io/ioutil"
	"math/rand"
)

func ExamplePolyline() {
	rand.Seed(0)
	data := make([]int, 1000)
	for i := range data {
		data[i] = (i + 1) / 200 * 200
	}
	for i := 0; i < 50; i++ {
		data[rand.Intn(800)+200] = rand.Intn(100)
	}
	ioutil.WriteFile("./_doc/demo.svg", []byte(Polyline(data)), 0644)
	// ioutil.WriteFile("./_doc/demo.svg", []byte(Polyline([]int{-10, -10, -10, 0, 0, 0, 10, 10, 10, 20, 20, 20, 30, 30, 30})), 0644)
	fmt.Println("ok")

	// Output:
	// ok
}
