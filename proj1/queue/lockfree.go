package queue

import (
	"sync/atomic"
	"unsafe"
)

type Request struct {
	Command   string  `json:"command"`
	Id        int     `json:"id"`
	Body      string  `json:"body"`
	TimeStamp float64 `json:"timestamp"`
}

// LockfreeQueue represents a FIFO structure with operations to enqueue
// and dequeue tasks represented as Request
type LockFreeQueue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
	size int
}

type node struct {
	value Request
	next  unsafe.Pointer
}

// NewQueue creates and initializes a LockFreeQueue
func NewLockFreeQueue() *LockFreeQueue {
	n := unsafe.Pointer(&node{})
	return &LockFreeQueue{head: n, tail: n, size: 0}
}

// Enqueue adds a series of Request to the queue
func (queue *LockFreeQueue) Enqueue(task *Request) {
	n := &node{value: *task}
	for {
		tail := load(&queue.tail)
		next := load(&tail.next)
		if tail == load(&queue.tail) { // are tail and next consistent?
			if next == nil {
				if cas(&tail.next, next, n) {
					cas(&queue.tail, tail, n)
					queue.size++ // Enqueue is done.  try to swing tail to the inserted node
					return
				}
			} else { // tail was not pointing to the last node
				// try to swing Tail to the next node
				cas(&queue.tail, tail, next)
				queue.size++
			}
		}
	}
}

// Dequeue removes a Request from the queue
func (queue *LockFreeQueue) Dequeue() *Request {
	for {
		head := load(&queue.head)
		tail := load(&queue.tail)
		next := load(&head.next)

		if head == load(&queue.head) {
			if head == tail {
				if next == nil {
					return nil
				}
				cas(&queue.tail, tail, next)
				queue.size--
			} else {
				v := next.value
				if cas(&queue.head, head, next) {
					queue.size--
					return &v
				}
			}
		}
	}

}

func load(p *unsafe.Pointer) (n *node) {
	return (*node)(atomic.LoadPointer(p))
}

func cas(p *unsafe.Pointer, old, new *node) (ok bool) {
	return atomic.CompareAndSwapPointer(
		p, unsafe.Pointer(old), unsafe.Pointer(new))
}
