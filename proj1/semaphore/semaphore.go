package semaphore

import "sync"

type Semaphore struct {
	capacity int
	cond     sync.Cond
}

func NewSemaphore(capacity int) *Semaphore {
	return &Semaphore{capacity: capacity, cond: *sync.NewCond(new(sync.Mutex))}
}

func (s *Semaphore) Up() {
	s.cond.L.Lock()

	s.capacity += 1
	s.cond.Signal()

	s.cond.L.Unlock()

}

func (s *Semaphore) Down() {
	s.cond.L.Lock()

	for s.capacity == 0 {
		s.cond.Wait()
	}

	s.capacity -= 1
	s.cond.L.Unlock()
}
