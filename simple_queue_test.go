package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleQueueLen(t *testing.T) {
	t.Parallel()
	q := NewSimpleQueue()
	assert.Equal(t, 0, q.Len())
}

func TestSimpleQueueCap(t *testing.T) {
	t.Parallel()
	q := NewSimpleQueue()
	assert.Equal(t, 0, q.Cap())
}

func TestSimpleQueueEnlarge(t *testing.T) {
	t.Parallel()
	q := NewSimpleQueue()
	for i := 0; i < 3; i++ {
		q.Push(i)
	}
	assert.Equal(t, 4, q.Cap())
}

func TestSimpleQueuePushPop1(t *testing.T) {
	t.Parallel()
	q := NewSimpleQueue()
	q.Push(1)
	q.Push(2)
	v := q.Pop()
	assert.Equal(t, 1, v.(int))
	q.Push("ss")
	assert.Equal(t, 2, q.Top())
	q.Push("q")
	assert.Equal(t, 2, q.Top())
	v = q.Pop()
	assert.Equal(t, 2, v.(int))
	v = q.Pop()
	assert.Equal(t, "ss", v.(string))
	v = q.Pop()
	assert.Equal(t, "q", v.(string))
	v = q.Pop()
	assert.Nil(t, v)
}

func TestSimpleQueuePushPop2(t *testing.T) {
	t.Parallel()
	q := NewSimpleQueue()
	for i := 0; i < 5; i++ {
		q.Push(i)
	}
	for i := 5; i < 15; i++ {
		v := q.Pop()
		assert.Equal(t, i-5, v.(int))
		q.Push(i)
	}
	assert.Equal(t, true, q.Push(15))
	for i := 16; i < 21; i++ {
		q.Push(i)
	}
	assert.Equal(t, &simpleQueue{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, q)
	assert.Equal(t, 10, q.Top())
}
