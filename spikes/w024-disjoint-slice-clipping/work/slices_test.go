package slices

import (
	"sync"
	"testing"
)

func TestSplitDisjointDataRaceFree(t *testing.T) {
	data := make([]int, 1000)
	for i := range data {
		data[i] = i
	}

	left, right := SplitDisjoint(data, 500)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := range left {
			left[i] += 1
		}
	}()

	go func() {
		defer wg.Done()
		for i := range right {
			right[i] += 2
		}
	}()

	wg.Wait()

	if left[0] != 1 || right[0] != 502 {
		t.Fatalf("unexpected mutated values: left[0]=%d, right[0]=%d", left[0], right[0])
	}
}

func BenchmarkSplitDisjoint(b *testing.B) {
	data := make([]int, 1024)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		l, r := SplitDisjoint(data, 512)
		_ = l
		_ = r
	}
}
