package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLen(t *testing.T) {
	t.Parallel()
	q := NewQueue(10)
	assert.Equal(t, 0, q.Len())
}

func TestCap(t *testing.T) {
	t.Parallel()
	q := NewQueue(10)
	assert.Equal(t, 10, q.Cap())
}

func TestEnlarge(t *testing.T) {
	t.Parallel()
	q := NewQueue(2)
	for i := 0; i < 3; i++ {
		q.Push(i)
	}
	assert.Equal(t, 4, q.Cap())
}

func TestPushPop1(t *testing.T) {
	t.Parallel()
	q := NewQueue(2)
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

func TestPushPop2(t *testing.T) {
	t.Parallel()
	q := queue{cap: 10, head: -1,
		data:  make([]interface{}, 10),
		limit: 1<<32 - 1}
	for i := 0; i < 5; i++ {
		q.Push(i)
	}
	for i := 5; i < 15; i++ {
		v := q.Pop()
		assert.Equal(t, i-5, v.(int))
		q.Push(i)
		if i == 7 {
			assert.Equal(t, 7, q.head)
			assert.Equal(t, 3, q.tail)
		}
		if i == 11 {
			assert.Equal(t, 1, q.head)
			assert.Equal(t, 7, q.tail)
		}
	}
	assert.Equal(t, []interface{}{10, 11, 12, 13, 14, 5, 6, 7, 8, 9}, q.data)
	assert.Equal(t, 4, q.head)
	assert.Equal(t, 0, q.tail)
	assert.Equal(t, true, q.Push(15))
	assert.Equal(t, []interface{}{10, 11, 12, 13, 14, 15, 6, 7, 8, 9}, q.data)
	for i := 16; i < 21; i++ {
		q.Push(i)
	}
	assert.Equal(t, []interface{}{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, interface{}(nil), interface{}(nil), interface{}(nil), interface{}(nil), interface{}(nil), interface{}(nil), interface{}(nil), interface{}(nil), interface{}(nil)}, q.data)
	assert.Equal(t, 20, q.cap)
	assert.Equal(t, 11, q.len)
	assert.Equal(t, 10, q.head)
	assert.Equal(t, 0, q.tail)
	assert.Equal(t, 10, q.Top())
}

func TestError(t *testing.T) {
	q := NewQueueWithLimit(1, 1)
	q.Push(1)
	assert.Equal(t, false, q.Push(2))
	assert.Equal(t, ErrHitLimit, q.Err())
	q = NewQueueWithLimit(2, 1)
	q.Push(1)
	assert.Equal(t, true, q.Push(2))
	assert.Nil(t, q.Err())
	q.Push(12)
	assert.Equal(t, ErrHitLimit, q.Err())
	q = NewQueueWithLimit(2, 3)
	for i := 0; i < 4; i++ {
		q.Push(i)
	}
	assert.Equal(t, ErrHitLimit, q.Err())
}
