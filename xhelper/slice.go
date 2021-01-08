package xhelper

import (
	"sync"
)

// SyncSlice for concurrent access
type SyncSlice struct {
	sync.RWMutex
	items []interface{}
}

// SyncSliceItem is an element in the slice
type SyncSliceItem struct {
	Idx int
	Val interface{}
}

// NewSyncSlice constructs a concurrent slice
func NewSyncSlice() *SyncSlice {
	return &SyncSlice{
		items: make([]interface{}, 0),
	}
}

// Append adds an element to the slice
func (s *SyncSlice) Append(item interface{}) {
	s.Lock()
	defer s.Unlock()
	s.items = append(s.items, item)
}

// Count returns the length of all elements in the slice
func (s *SyncSlice) Count() int {
	s.Lock()
	defer s.Unlock()
	return len(s.items)
}

// Get returns the element at index of slice or nil
func (s *SyncSlice) Get(i int) <-chan SyncSliceItem {
	var ch = make(chan SyncSliceItem)
	go func() {
		s.Lock()
		defer s.Unlock()
		ch <- SyncSliceItem{i, s.items[i]} // todo be careful with out of range errors - use Count
		close(ch)
	}()

	return ch
}

// Iter outputs each item over a channel to the caller
func (s *SyncSlice) Iter() <-chan SyncSliceItem {
	var ch = make(chan SyncSliceItem)

	go func() {
		s.Lock()
		defer s.Unlock()
		for i, val := range s.items {
			ch <- SyncSliceItem{i, val}
		}
		close(ch)
	}()

	return ch
}
