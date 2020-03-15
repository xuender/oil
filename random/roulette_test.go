package random

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestInt struct {
	score int
}

func (t TestInt) Score() int {
	return t.score
}

func TestRoulette_Take(t *testing.T) {
	rand.Seed(3)
	r := NewRoulette([]Scorer{TestInt{score: 2}, TestInt{score: 1}, TestInt{score: 7}})
	assert.Equal(t, 7, r.Take().(TestInt).score, "0")
	assert.Equal(t, 7, r.Take().(TestInt).score, "1")
	assert.Equal(t, 7, r.Take().(TestInt).score, "2")
	assert.Equal(t, 1, r.Take().(TestInt).score, "3")
	assert.Equal(t, 7, r.Take().(TestInt).score, "4")
	assert.Equal(t, 2, r.Take().(TestInt).score, "5")
	assert.Equal(t, 7, r.Take().(TestInt).score, "6")
	assert.Equal(t, 7, r.Take().(TestInt).score, "7")
	assert.Equal(t, 7, r.Take().(TestInt).score, "8")
	assert.Equal(t, 7, r.Take().(TestInt).score, "9")
}

// 轮盘的原理
func ExampleNewRoulette() {
	arr := []int{2, 2, 6}
	add := 0
	for i := range arr {
		arr[i] = arr[i] + add
		add = arr[i]
	}
	fmt.Println(arr)

	// 计算比例
	m := map[int]int{2: 0, 4: 0, 10: 0}
	for i := 0; i < 10000; i++ {
		r := rand.Intn(10)
		if r < 2 {
			m[2]++
		} else if r < 4 {
			m[4]++
		} else {
			m[10]++
		}
	}
	fmt.Println(math.Round(float64(m[2])/1000), math.Round(float64(m[4])/1000), math.Round(float64(m[10])/1000))

	// Output:
	// [2 4 10]
	// 2 2 6
}
