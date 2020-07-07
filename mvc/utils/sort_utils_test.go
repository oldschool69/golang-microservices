package utils

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestBubbleSortWorstCase(t *testing.T) {
	// Initialization
	els := []int{9,8,7,6,5}
	sortedEls := []int{5,6,7,8,9}

	// Execution
	els = BubbleSort(els)

	// Validation
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, sortedEls, els)
}

func TestBubbleSortBestCase(t *testing.T) {
	// Initialization
	els := []int{5,6,7,8,9}
	sortedEls := []int{5,6,7,8,9}

	// Execution
	els = BubbleSort(els)

	// Validation
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, sortedEls, els)
}

func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n-1; j >= 0; j-- {
		result[i] = j
		i++
	}

	return result
}

func TestGetElements(t *testing.T) {
	els := getElements(5)
	resultEls := []int{4,3,2,1,0}
	assert.NotNil(t, els)
	assert.EqualValues(t, resultEls, els)
}

func BenchmarkBubbleSort10(b *testing.B) {
	els := getElements(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkSort10(b *testing.B) {
	els := getElements(10)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}


func BenchmarkBubbleSort1000(b *testing.B) {
	els := getElements(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkSort1000(b *testing.B) {
	els := getElements(1000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	els := getElements(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkSort100000(b *testing.B) {
	els := getElements(100000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}