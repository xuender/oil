package random

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	rand.Seed(0)
	q := NewQueue(10, 3)
	assert.Equal(t, 3, len(q), "队列长度")
	assert.Contains(t, q, 3, "3")
	assert.Contains(t, q, 4, "4")
	assert.Contains(t, q, 6, "6")
}
func ExampleNewQueue() {
	rand.Seed(0)
	q := NewQueue(10, 3)
	sort.Ints(q)
	fmt.Println(q)

	// Output:
	// [3 4 6]
}
