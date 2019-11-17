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
	Add(interface{})
	Remove(interface{}) bool
	Contains(interface{}) bool
	Replace(interface{}, interface{}) bool
	Slice() []interface{}
}

type set struct {
	sync.RWMutex
	m   map[interface{}]int32
	len int32
}

// New declares a new set.
func New() Set {
	return &set{
		m: make(map[interface{}]int32),
	}
}

// NewFromSlice declares a new set from a
// supplied slice.
func NewFromSlice(slc []interface{}) Set {
	s := New()
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

// Add a value to the set.
func (s *set) Add(value interface{}) {
	s.Lock()
	if _, loaded := s.m[value]; !loaded {
		s.m[value] = s.len
		atomic.AddInt32(&s.len, 1)
	}
	s.Unlock()
}

// Remove a value from the set.
func (s *set) Remove(value interface{}) (removed bool) {
	s.Lock()
	_, removed = s.m[value]
	delete(s.m, value)
	atomic.CompareAndSwapInt32(&s.len, s.len, s.len-1)
	s.Unlock()
	return
}

// Contains returns true if the supplied value
// exists in the set, otherwise false.
func (s *set) Contains(value interface{}) (contained bool) {
	s.RLock()
	_, contained = s.m[value]
	s.RUnlock()
	return
}

// Replace a value in the set
func (s *set) Replace(old, new interface{}) (replaced bool) {
	s.Lock()
	if v, loaded := s.m[old]; loaded {
		delete(s.m, old)
		s.m[new] = v
		replaced = true
	}
	s.Unlock()
	return
}

// Slice returns the values held in the set
// in the form of a Go slice object.
func (s *set) Slice() (slc []interface{}) {
	s.RLock()
	slc = make([]interface{}, s.len)

	for k, v := range s.m {
		slc[v] = k
	}

	s.RUnlock()
	return
}
