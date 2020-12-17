package queue

import (
	"errors"
)

// Queue is the FIFO queue
type Queue interface {
	Push(v interface{}) bool // Push v into the queue
	Pop() interface{}        // Pop the first element from queue
	Err() error              // Return err when failed
	Len() int                // Return current element count of queue
	Cap() int                // Return current capacity of queue
	Top() interface{}        // Return the oldest element that will be popped if Pop() is called
}

type queue struct {
	head  int
	tail  int
	cap   int
	len   int
	limit int
	err   error
	data  []interface{}
}

// ErrHitLimit is the error when limit hit
var ErrHitLimit = errors.New("hit the queue limit")

// NewQueue creates a new queue without limit
func NewQueue(cap int) Queue {
	return &queue{
		cap:   cap,
		head:  -1,
		data:  make([]interface{}, cap),
		limit: int(^uint(0) >> 1), // max int
	}
}

// NewQueueWithLimit create a queue with limit
func NewQueueWithLimit(cap, limit int) Queue {
	if limit < cap {
		limit = cap
	}
	return &queue{
		cap:   cap,
		head:  -1,
		data:  make([]interface{}, cap),
		limit: limit,
	}
}

func (q *queue) Len() int { return q.len }

func (q *queue) Cap() int { return q.cap }

func (q *queue) enlarge() bool {
	if q.cap+1 >= q.limit {
		q.err = ErrHitLimit
		return false
	}
	newCap := q.cap * 2
	if newCap > q.limit {
		newCap = q.limit
	}
	d := make([]interface{}, newCap)
	if q.head < q.tail {
		l1 := len(q.data) - q.tail
		copy(d[0:l1], q.data[q.tail:len(q.data)])
		copy(d[l1:l1+q.head+1], q.data[0:q.head+1])
		q.head = l1 + q.head
	} else {
		copy(d[0:q.head+1-q.tail], q.data[q.tail:q.head+1])
		q.head = q.head - q.tail
	}
	q.tail = 0
	q.cap = newCap
	q.data = d
	return true
}

func (q *queue) Push(v interface{}) bool {
	if q.len >= q.cap {
		if !q.enlarge() {
			return false
		}
	}
	q.head++
	if q.head >= q.cap {
		q.head = 0
	}
	q.len++
	q.data[q.head] = v
	return true
}

func (q *queue) Pop() interface{} {
	if q.len <= 0 {
		return nil
	}
	v := q.data[q.tail]
	q.tail++
	if q.tail >= q.cap {
		q.tail = 0
	}
	q.len--
	return v
}

func (q *queue) Top() interface{} {

	return q.data[q.tail]
}

func (q *queue) Err() error {
	return q.err
}
