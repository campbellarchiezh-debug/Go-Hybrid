package session

import (
	"testing"
)

func BenchmarkSessionProtocolThroughput(b *testing.B) {
	ch := make(chan int, 1)
	done := make(chan struct{})

	go func() {
		for i := 0; i < b.N; i++ {
			msg := <-ch
			ch <- msg + 1
		}
		close(done)
	}()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ch <- i
		<-ch
	}

	<-done
}

func TestSessionProtocolOrdering(t *testing.T) {
	sc := NewSessionChan[PingState, int](2)
	ch := sc.BareChannel()

	ch <- 42
	val := <-ch
	if val != 42 {
		t.Fatalf("expected 42, got %d", val)
	}
}
