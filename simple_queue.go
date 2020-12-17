package queue

// simpleQueue the simple slice queue
type simpleQueue []interface{}

// NewSimpleQueue return the simple slice queue
func NewSimpleQueue() Queue {
	return &simpleQueue{}
}

func (q *simpleQueue) Len() int { return len(*q) }

func (q *simpleQueue) Cap() int { return cap(*q) }

func (q *simpleQueue) Push(v interface{}) bool {
	*q = append(*q, v)
	return true
}

func (q *simpleQueue) Pop() interface{} {
	if len(*q) < 1 {
		return nil
	}
	v := (*q)[0]
	*q = append((*q)[:0:0], (*q)[1:len(*q)]...)
	return v
}

func (q *simpleQueue) Err() error {
	return nil
}

func (q *simpleQueue) Top() interface{} {
	if len(*q) < 1 {
		return nil
	}
	return (*q)[0]
}
