package rbtree_test

import (
	"fmt"
	"sort"
	"testing"
	"unsafe"

	"github.com/cdongyang/rbtree"
)

type IntSetNode struct {
	n rbtree.SetNode
}

func (n IntSetNode) GetData() int {
	return n.n.GetData().(int)
}

func (n IntSetNode) Next() IntSetNode {
	return IntSetNode{n: n.n.Next()}
}

func (n IntSetNode) Last() IntSetNode {
	return IntSetNode{n: n.n.Next()}
}

func (n IntSetNode) GetSet() *IntSet {
	return (*IntSet)(unsafe.Pointer(n.GetSet()))
}

type IntSet struct {
	set rbtree.Set
}

func NewIntSet(compare func(a, b int) int) *IntSet {
	var s = &IntSet{}
	s.Init(true, compare)
	return s
}

func NewMultiIntSet(compare func(a, b int) int) *IntSet {
	var s = &IntSet{}
	s.Init(false, compare)
	return s
}

func (s *IntSet) pack(n rbtree.SetNode) IntSetNode {
	return IntSetNode{n: n}
}

func (s *IntSet) Init(unique bool, compare func(a, b int) int) {
	s.set.Init(unique, int(0), func(a, b interface{}) int {
		return compare(a.(int), b.(int))
	})
}

func (s *IntSet) Begin() IntSetNode {
	return s.pack(s.set.Begin())
}

func (s *IntSet) End() IntSetNode {
	return s.pack(s.set.End())
}

func (s *IntSet) EqualRange(data int) (beg, end IntSetNode) {
	a, b := s.set.EqualRange(data)
	return s.pack(a), s.pack(b)
}

func (s *IntSet) EraseNode(n IntSetNode) {
	s.set.EraseNode(n.n)
}

func (s *IntSet) EraseNodeRange(beg, end IntSetNode) (count int) {
	return s.set.EraseNodeRange(beg.n, end.n)
}

func (s *IntSet) Find(data int) IntSetNode {
	return s.pack(s.set.Find(data))
}

func (s *IntSet) Insert(data int) (IntSetNode, bool) {
	// if data type is not direct interface, use NoescapeInterface to avoid data escape to heap,
	// thus reduce heap objects.
	// but if you don't know wheather the data type is direct interface, don't use NoescapeInterface !!!
	n, ok := s.set.Insert(rbtree.NoescapeInterface(data))
	return s.pack(n), ok
}

func (s *IntSet) LowerBound(data int) IntSetNode {
	return s.pack(s.set.LowerBound(data))
}

func (s *IntSet) UpperBound(data int) IntSetNode {
	return s.pack(s.set.UpperBound(data))
}

func (s *IntSet) Count(data int) (count int) {
	return s.set.Count(data)
}

func (s *IntSet) Erase(data int) (count int) {
	return s.set.Erase(data)
}

func (s *IntSet) Size() int {
	return s.set.Size()
}

func TestIntSet(t *testing.T) {
	t.Run("method", func(t *testing.T) {
		var slice = []int{1, 4, 6, 5, 3, 7, 2, 9}
		var cpSlice = make([]int, len(slice))
		copy(cpSlice, slice)
		sort.IntSlice(cpSlice).Sort()
		s := NewIntSet(func(a, b int) int {
			return a - b
		})
		for _, val := range slice {
			s.Insert(val)
		}
		for it, i := s.Begin(), 0; it != s.End(); it = it.Next() {
			if it.GetData() != cpSlice[i] {
				t.Fatal(it.GetData(), cpSlice[i])
			}
			i++
		}
		for it := s.Begin(); it != s.End(); {
			tmp := it.GetData()
			it = it.Next()
			s.Erase(tmp)
		}
		if s.Size() != 0 {
			t.Fatal(s.Size())
		}
	})
	t.Run("escape", func(t *testing.T) {
		var x = 1
		s := NewIntSet(func(a, b int) int { return a - b })
		n := testing.AllocsPerRun(1000, func() {
			var tmp = x
			s.Insert(tmp)
			x++
		})
		if n > 0 {
			t.Fatal("insert escape", n)
		}
		n = testing.AllocsPerRun(1000, func() {
			var tmp = 10
			s.Find(tmp)
			x++
		})
		if n > 0 {
			t.Fatal("find escape", n)
		}
		n = testing.AllocsPerRun(1000, func() {
			var tmp = 10
			s.Erase(tmp)
			x++
		})
		if n > 0 {
			t.Fatal("erase escape", n)
		}
	})
}

type intSet struct { // hide node
	set rbtree.Set
}

func NewintSet(compare func(a, b int) int) *intSet {
	var s = &intSet{}
	s.set.Init(true, int(0), func(a, b interface{}) int {
		return compare(a.(int), b.(int))
	})
	return s
}

func (s *intSet) Insert(data int) (ok bool) {
	_, ok = s.set.Insert(rbtree.NoescapeInterface(data))
	return ok
}

func (s *intSet) Find(data int) (ok bool) {
	return s.set.Find(data) == s.set.End()
}

func (s *intSet) Erase(data int) (ok bool) {
	return s.set.Erase(data) == 1
}

func (s *intSet) Range(f func(data int) bool) {
	for it := s.set.Begin(); it != s.set.End(); it = it.Next() {
		ok := f(it.GetData().(int))
		if !ok {
			return
		}
	}
}

func (s *intSet) Size() int {
	return s.set.Size()
}

func TestIntSetNoIterator(t *testing.T) {
	t.Run("method", func(t *testing.T) {
		var slice = []int{1, 4, 6, 5, 3, 7, 2, 9}
		var cpSlice = make([]int, len(slice))
		copy(cpSlice, slice)
		sort.IntSlice(cpSlice).Sort()
		s := NewintSet(func(a, b int) int {
			return a - b
		})
		for _, val := range slice {
			s.Insert(val)
		}
		i := 0
		s.Range(func(data int) bool {
			if data != cpSlice[i] {
				t.Fatal(data, cpSlice[i])
				return false
			}
			i++
			return true
		})
		s.Range(func(data int) bool {
			if !s.Erase(data) {
				t.Fatal(data)
				return false
			}
			return true
		})
		if s.Size() != 0 {
			t.Fatal(s.Size())
		}
	})
	t.Run("escape", func(t *testing.T) {
		var x = 1
		s := NewintSet(func(a, b int) int { return a - b })
		n := testing.AllocsPerRun(1000, func() {
			var tmp = x
			s.Insert(tmp)
			x++
		})
		if n > 0 {
			t.Fatal("insert escape", n)
		}
		n = testing.AllocsPerRun(1000, func() {
			var tmp = 10
			s.Find(tmp)
			x++
		})
		if n > 0 {
			t.Fatal("find escape", n)
		}
		n = testing.AllocsPerRun(1000, func() {
			var tmp = 10
			s.Erase(tmp)
			x++
		})
		if n > 0 {
			t.Fatal("erase escape", n)
		}
	})
}

type intMap struct {
	mp rbtree.Map
}

func NewIntMap(compare func(a, b int) int) *intMap {
	var m = &intMap{}
	m.mp.Init(true, int(0), new(int), func(a, b interface{}) int {
		return compare(a.(int), b.(int))
	})
	return m
}

func (m *intMap) StoreOrLoad(key int, val *int) (ok bool, actual *int) {
	n, ok := m.mp.Insert(rbtree.NoescapeInterface(key), val)
	return ok, n.GetVal().(*int)
}

func (m *intMap) Load(key int) (ok bool, val *int) {
	n := m.mp.Find(key)
	if n == m.mp.End() {
		return false, nil
	}
	return true, n.GetVal().(*int)
}

func (m *intMap) Delete(key int) (ok bool) {
	return m.mp.Erase(key) == 1
}

func (m *intMap) Range(f func(key int, val *int) bool) {
	for it := m.mp.Begin(); it != m.mp.End(); it = it.Next() {
		ok := f(it.GetKey().(int), it.GetVal().(*int))
		if !ok {
			return
		}
	}
}

func (m *intMap) Size() int {
	return m.mp.Size()
}
func TestIntMapNoIterator(t *testing.T) {
	t.Run("method", func(t *testing.T) {
		var slice = []int{1, 4, 6, 5, 3, 7, 2, 9}
		var cpSlice = make([]int, len(slice))
		copy(cpSlice, slice)
		sort.IntSlice(cpSlice).Sort()
		s := NewIntMap(func(a, b int) int {
			return a - b
		})
		for i := range slice {
			s.StoreOrLoad(slice[i], &slice[i])
		}
		i := 0
		s.Range(func(key int, val *int) bool {
			if key != cpSlice[i] || *val != cpSlice[i] {
				t.Fatal(key, *val, cpSlice[i])
				return false
			}
			i++
			return true
		})
		s.Range(func(key int, val *int) bool {
			if !s.Delete(key) {
				t.Fatal(key, *val)
				return false
			}
			return true
		})
		if s.Size() != 0 {
			t.Fatal(s.Size())
		}
	})
	t.Run("escape", func(t *testing.T) {
		var x = 1
		s := NewIntMap(func(a, b int) int { return a - b })
		n := testing.AllocsPerRun(1000, func() {
			var tmp = x
			s.StoreOrLoad(tmp, nil)
			x++
		})
		if n > 0 {
			t.Fatal("insert escape", n)
		}
		n = testing.AllocsPerRun(1000, func() {
			var tmp = 10
			s.Load(tmp)
			x++
		})
		if n > 0 {
			t.Fatal("find escape", n)
		}
		n = testing.AllocsPerRun(1000, func() {
			var tmp = 10
			s.Delete(tmp)
			x++
		})
		if n > 0 {
			t.Fatal("erase escape", n)
		}
	})
}

func ExampleMap() {
	var slice = []int{1, 4, 6, 5, 3, 7, 2, 9}
	// key type: int, value type: *int
	mp := rbtree.NewMap(int(0), new(int), func(a, b interface{}) int {
		return a.(int) - b.(int)
	})
	for i := range slice {
		mp.Insert(slice[i], &slice[i])
	}
	var indexOf = func(p *int) int {
		for i := range slice {
			if &slice[i] == p {
				return i
			}
		}
		return -1
	}
	// iterator
	for it, i := mp.Begin(), 0; it != mp.End(); it = it.Next() {
		fmt.Println(it.GetKey(), indexOf(it.GetVal().(*int)))
		i++
	}
	mp = nil // free tree to make it collect by GC
	//Output:
	//1 0
	//2 6
	//3 4
	//4 1
	//5 3
	//6 2
	//7 5
	//9 7
}
