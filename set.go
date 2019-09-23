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
	Len() uint64
	Add(interface{})
	Remove(interface{})
	Contains(interface{}) bool
	Slice() []interface{}
}

type set struct {
	m   sync.Map
	len uint64
}

// New declares a new set.
func New() Set {
	return &set{
		m: sync.Map{},
	}
}

// Len returns the number of values stored
// in the set.
func (s *set) Len() uint64 {
	return s.len
}

// Add a value to the set.
func (s *set) Add(value interface{}) {
	if _, loaded := s.m.LoadOrStore(value, struct{}{}); !loaded {
		atomic.AddUint64(&s.len, 1)
	}
}

// Remove a value from the set.
func (s *set) Remove(value interface{}) {
	defer atomic.CompareAndSwapUint64(&s.len, s.len, s.len-1)
	s.m.Delete(value)
}

// Contains returns true if the supplied value
// exists in the set, otherwise false.
func (s *set) Contains(value interface{}) (contained bool) {
	_, contained = s.m.Load(value)
	return
}

// Slice returns the values held in the set
// in the form of a Go slice object.
func (s *set) Slice() (slc []interface{}) {
	slc = make([]interface{}, 0, s.len)

	s.m.Range(func(key, value interface{}) bool {
		slc = append(slc, key)
		return true
	})

	return
}
