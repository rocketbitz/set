package set

import (
	"sync"
	"sync/atomic"
)

// Set defines an interface for a unique set
// object which uses a sync.Pool and atomic
// counter as its underlying data structures.
// Due to the use of these underlying objects,
// the Set is safe to be accessed by multiple
// go-routines, but all limitations of the
// sync.Pool object apply to Set.
type Set interface {
	Len() int
	At(int) interface{}
	Add(interface{})
	Remove(interface{}) bool
	Contains(interface{}) bool
	Replace(interface{}, interface{}) bool
	Index(interface{}) int
	Slice() []interface{}
}

type set struct {
	sync.RWMutex
	m    map[interface{}]int32
	len  int32
	safe bool
}

// New declares a new thread-safe set.
func New() Set {
	return &set{
		m:    make(map[interface{}]int32),
		safe: true,
	}
}

// NewUnsafe declares a new thread-unsafe set.
func NewUnsafe() Set {
	return &set{m: make(map[interface{}]int32)}
}

// NewFromSlice declares a new thread-safe set
// from a supplied slice.
func NewFromSlice(slc []interface{}) Set {
	s := New()
	for _, v := range slc {
		s.Add(v)
	}
	return s
}

// NewUnsafeFromSlice declares a new thread-
// unsafe set from a supplied slice.
func NewUnsafeFromSlice(slc []interface{}) Set {
	s := NewUnsafe()
	for _, v := range slc {
		s.Add(v)
	}
	return s
}

// Len returns the number of values stored
// in the set.
func (s *set) Len() int {
	return int(s.len)
}

// At retrives the value in the set at the
// supplied index.
func (s *set) At(i int) interface{} {
	for k, v := range s.m {
		if int(v) == i {
			return k
		}
	}
	return nil
}

// Add a value to the set.
func (s *set) Add(value interface{}) {
	if s.safe {
		s.Lock()
		defer s.Unlock()
	}

	if _, loaded := s.m[value]; !loaded {
		s.m[value] = s.len
		atomic.AddInt32(&s.len, 1)
	}
}

// Remove a value from the set.
func (s *set) Remove(value interface{}) (removed bool) {
	if s.safe {
		s.Lock()
		defer s.Unlock()
	}

	_, removed = s.m[value]
	delete(s.m, value)
	atomic.CompareAndSwapInt32(&s.len, s.len, s.len-1)

	return
}

// Contains returns true if the supplied value
// exists in the set, otherwise false.
func (s *set) Contains(value interface{}) (contained bool) {
	if s.safe {
		s.RLock()
		defer s.RUnlock()
	}

	_, contained = s.m[value]

	return
}

// Index of the specified value in the set. -1
// if the set does not contain the value
func (s *set) Index(value interface{}) int {
	if s.safe {
		s.RLock()
		defer s.RUnlock()
	}

	i, ok := s.m[value]
	if ok {
		return int(i)
	}

	return -1
}

// Replace a value in the set
func (s *set) Replace(old, new interface{}) (replaced bool) {
	if s.safe {
		s.Lock()
		defer s.Unlock()
	}

	if v, loaded := s.m[old]; loaded {
		delete(s.m, old)
		s.m[new] = v
		replaced = true
	}

	return
}

// Slice returns the values held in the set
// in the form of a Go slice object.
func (s *set) Slice() (slc []interface{}) {
	if s.safe {
		s.RLock()
		defer s.RUnlock()
	}

	slc = make([]interface{}, s.len)

	for k, v := range s.m {
		slc[v] = k
	}

	return
}
