// Package lock provides an implementation of a read-write lock
// that uses condition variables and mutexes.
package lock

import "sync"

type RWLock interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
}

type RWCondRWLock struct {
	readerCount int
	cond        *sync.Cond
}

func (l *RWCondRWLock) RLock() {
	l.cond.L.Lock()

	for l.readerCount == 32 {
		l.cond.Wait()
	}
	l.readerCount++
	l.cond.L.Unlock()
}

func (l *RWCondRWLock) RUnlock() {
	l.cond.L.Lock()
	l.readerCount--
	if l.readerCount == 0 {
		l.cond.Signal()
	}
	l.cond.L.Unlock()
}

func (l *RWCondRWLock) Lock() {
	l.cond.L.Lock()
	for l.readerCount > 0 {
		l.cond.Wait()
	}
}

func (l *RWCondRWLock) Unlock() {
	l.cond.Signal()
	l.cond.L.Unlock()
}

func NewRWLock() *RWCondRWLock {
	return &RWCondRWLock{0, sync.NewCond(new(sync.Mutex))}
}
