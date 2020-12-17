package queue

import (
	"container/list"
	"math/rand"
	"testing"
)

const size = 2048

func BenchmarkPushQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var q Queue = NewQueue(size)
		for n := 0; n < size; n++ {
			q.Push(n)
		}
	}
}

func BenchmarkPushSimpleQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var q Queue = NewSimpleQueue()
		for n := 0; n < size; n++ {
			q.Push(n)
		}
	}
}

func BenchmarkPushList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var q list.List
		for n := 0; n < size; n++ {
			q.PushFront(n)
		}
	}
}

func BenchmarkPushChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := make(chan interface{}, size)
		for n := 0; n < size; n++ {
			q <- n
		}
		close(q)
	}
}

var rands []float32

func makeRands() {
	if rands != nil {
		return
	}
	rand.Seed(64738)
	for i := 0; i < 2*size; i++ {
		rands = append(rands, rand.Float32())
	}
}
func BenchmarkRandomQueue(b *testing.B) {
	makeRands()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var q Queue = NewQueue(size)
		for n := 0; n < 2*size; n += 2 {
			if rands[n] < 0.8 {
				q.Push(n)
			}
			if rands[n+1] < 0.5 {
				q.Pop()
			}
		}
	}
}

func BenchmarkRandomSimpleQueue(b *testing.B) {
	makeRands()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var q Queue = NewSimpleQueue()
		for n := 0; n < 2*size; n += 2 {
			if rands[n] < 0.8 {
				q.Push(n)
			}
			if rands[n+1] < 0.5 {
				q.Pop()
			}
		}
	}
}
func BenchmarkRandomList(b *testing.B) {
	makeRands()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var q list.List
		for n := 0; n < 2*size; n += 2 {
			if rands[n] < 0.8 {
				q.PushFront(n)
			}
			if rands[n+1] < 0.5 {
				if e := q.Back(); e != nil {
					q.Remove(e)
				}
			}
		}
	}
}

func BenchmarkGrowShrinkQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var q Queue = NewQueue(2)
		for n := 0; n < size; n++ {
			q.Push(i)
		}
		for n := 0; n < size; n++ {
			q.Pop()
		}
	}
}

func BenchmarkGrowShrinkSimpleQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var q Queue = NewSimpleQueue()
		for n := 0; n < size; n++ {
			q.Push(i)
		}
		for n := 0; n < size; n++ {
			q.Pop()
		}
	}
}

func BenchmarkGrowShrinkList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var q list.List
		for n := 0; n < size; n++ {
			q.PushBack(i)
		}
		for n := 0; n < size; n++ {
			if e := q.Front(); e != nil {
				q.Remove(e)
			}
		}
	}
}
