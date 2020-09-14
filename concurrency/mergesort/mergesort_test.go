package mergesort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const size = 10000

//const size = 1000000

func TestMergesortSequential(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	mergeSortSequential(arr)

	assert.Equal(t, []int{1, 2, 3, 4, 5}, arr)
}

func TestMergesortParallel1(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	mergeSortParallel1(arr)

	assert.Equal(t, []int{1, 2, 3, 4, 5}, arr)
}

func BenchmarkMergesortSequential(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := random(size)
		b.StartTimer()
		mergeSortSequential(arr)
		b.StopTimer()
	}
}

func BenchmarkMergesortParallel1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := random(size)
		b.StartTimer()
		mergeSortParallel1(arr)
		b.StopTimer()
	}
}

func BenchmarkMergesortParallel2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := random(size)
		b.StartTimer()
		mergeSortParallel2(arr)
		b.StopTimer()
	}
}

func BenchmarkMergesortParallel3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := random(size)
		b.StartTimer()
		mergeSortParallel3(arr)
		b.StopTimer()
	}
}
