package goedbt

import "sync"

// modified version of https://github.com/gammazero/deque for use in the
// event-driven BT implementation. adds mutex for concurrent access

// minCapacity is the smallest capacity that deque may have. Must be power of 2
// for bitwise modulus: x % n == x & (n - 1).
const minCapacity = 4

// Deque represents a single instance of the deque data structure. A Deque
// instance contains items of the type specified by the type argument.
type Deque[T any] struct {
	buf    []T
	head   int
	tail   int
	count  int
	minCap int

	sync.Mutex
}

func NewDeque[T any](baseCap int) *Deque[T] {
	minCap := minCapacity
	// round baseCap up to the nearest power of 2
	for minCap < baseCap {
		minCap <<= 1
	}

	return &Deque[T]{
		minCap: minCap,
	}
}

// Cap returns the current capacity of the Deque. If q is nil, q.Cap() is zero.
func (q *Deque[T]) Cap() int {
	if q == nil {
		return 0
	}
	return len(q.buf)
}

// Len returns the number of elements currently stored in the queue. If q is
// nil, q.Len() returns zero.
func (q *Deque[T]) Len() int {
	if q == nil {
		return 0
	}
	return q.count
}

// PushBack appends an element to the back of the queue. Implements FIFO when
// elements are removed with PopFront, and LIFO when elements are removed with
// PopBack.
func (q *Deque[T]) PushBack(elem T) {
	q.Lock()
	defer q.Unlock()

	q.growIfFull()

	q.buf[q.tail] = elem
	// Calculate new tail position.
	q.tail = q.next(q.tail)
	q.count++
}

// PushFront prepends an element to the front of the queue.
func (q *Deque[T]) PushFront(elem T) {
	q.Lock()
	defer q.Unlock()

	q.growIfFull()

	// Calculate new head position.
	q.head = q.prev(q.head)
	q.buf[q.head] = elem
	q.count++
}

// PopFront removes and returns the element from the front of the queue.
// Implements FIFO when used with PushBack. If the queue is empty, the call
// panics.
func (q *Deque[T]) PopFront() T {
	q.Lock()
	defer q.Unlock()

	if q.count <= 0 {
		panic("deque: PopFront() called on empty queue")
	}
	ret := q.buf[q.head]
	var zero T
	q.buf[q.head] = zero
	// Calculate new head position.
	q.head = q.next(q.head)
	q.count--

	q.shrinkIfExcess()
	return ret
}

// PopBack removes and returns the element from the back of the queue.
// Implements LIFO when used with PushBack. If the queue is empty, the call
// panics.
func (q *Deque[T]) PopBack() T {
	q.Lock()
	defer q.Unlock()

	if q.count <= 0 {
		panic("deque: PopBack() called on empty queue")
	}

	// Calculate new tail position
	q.tail = q.prev(q.tail)

	// Remove value at tail.
	ret := q.buf[q.tail]
	var zero T
	q.buf[q.tail] = zero
	q.count--

	q.shrinkIfExcess()
	return ret
}

// Front returns the element at the front of the queue. This is the element
// that would be returned by PopFront. This call panics if the queue is empty.
func (q *Deque[T]) Front() T {
	q.Lock()
	defer q.Unlock()

	if q.count <= 0 {
		panic("deque: Front() called when empty")
	}
	return q.buf[q.head]
}

// Back returns the element at the back of the queue. This is the element that
// would be returned by PopBack. This call panics if the queue is empty.
func (q *Deque[T]) Back() T {
	q.Lock()
	defer q.Unlock()

	if q.count <= 0 {
		panic("deque: Back() called when empty")
	}
	return q.buf[q.prev(q.tail)]
}

// Clear removes all elements from the queue, but retains the current capacity.
// This is useful when repeatedly reusing the queue at high frequency to avoid
// GC during reuse. The queue will not be resized smaller as long as items are
// only added. Only when items are removed is the queue subject to getting
// resized smaller.
func (q *Deque[T]) Clear() {
	q.Lock()
	defer q.Unlock()

	var zero T
	modBits := len(q.buf) - 1
	h := q.head
	for i := 0; i < q.Len(); i++ {
		q.buf[(h+i)&modBits] = zero
	}
	q.head = 0
	q.tail = 0
	q.count = 0
}

// prev returns the previous buffer position wrapping around buffer.
func (q *Deque[T]) prev(i int) int {
	return (i - 1) & (len(q.buf) - 1) // bitwise modulus
}

// next returns the next buffer position wrapping around buffer.
func (q *Deque[T]) next(i int) int {
	return (i + 1) & (len(q.buf) - 1) // bitwise modulus
}

// growIfFull resizes up if the buffer is full.
func (q *Deque[T]) growIfFull() {
	if q.count != len(q.buf) {
		return
	}
	if len(q.buf) == 0 {
		if q.minCap == 0 {
			q.minCap = minCapacity
		}
		q.buf = make([]T, q.minCap)
		return
	}
	q.resize(q.count << 1)
}

// shrinkIfExcess resize down if the buffer 1/4 full.
func (q *Deque[T]) shrinkIfExcess() {
	if len(q.buf) > q.minCap && (q.count<<2) == len(q.buf) {
		q.resize(q.count << 1)
	}
}

// resize resizes the deque to fit exactly twice its current contents. This is
// used to grow the queue when it is full, and also to shrink it when it is
// only a quarter full.
func (q *Deque[T]) resize(newSize int) {
	newBuf := make([]T, newSize)
	if q.tail > q.head {
		copy(newBuf, q.buf[q.head:q.tail])
	} else {
		n := copy(newBuf, q.buf[q.head:])
		copy(newBuf[n:], q.buf[:q.tail])
	}

	q.head = 0
	q.tail = q.count
	q.buf = newBuf
}
