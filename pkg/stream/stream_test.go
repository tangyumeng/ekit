// Code generated by go2go; DO NOT EDIT.


//line stream_test.go2:1
package stream

//line stream_test.go2:1
import (
//line stream_test.go2:1
 "sort"
//line stream_test.go2:1
 "testing"
//line stream_test.go2:1
)

//line stream_test.go2:7
func TestStreamAPI(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	s := instantiate୦୦Of୦int(ints)

	if s == nil {
		t.Fail()
	}

	anyMatch := s.AnyMatch(func(e int) bool {
		return e == 3
	})

	if !anyMatch {
		t.Fatal("AnyMatch failed")
	}

	res := s.FindAny(func(e int) bool {
		return e == 4
	})

	if res == 0 {
		t.Fatal("FindAny failed")
	}

	allMatch := s.AllMatch(func(e int) bool {
		return e > 0
	})
	if !allMatch {
		t.Fatal("AllMatch failed")
	}

	ns := s.ConcatArray([]int{1, 2, 3})

	if ns.Count() != 8 {
		t.Fatal("ConCatArray Failed")
	}

	ds := s.Distinct(func(e1, e2 int) int {
		return e1 - e2
	})

	if ds.Count() != 5 {
		t.Fatal("Distinct Failed")
	}
}
//line stream.go2:13
func instantiate୦୦Of୦int(elems []int,) *instantiate୦୦stream୦int {
	return &instantiate୦୦stream୦int{
		eles: elems,
	}
}

//line stream.go2:17
type instantiate୦୦stream୦int struct {
//line stream.go2:21
 eles []int
	def int
}

func (s *instantiate୦୦stream୦int,) OrElse(e int,) *instantiate୦୦stream୦int {
	s.def = e
	return s
}

func (s *instantiate୦୦stream୦int,) Filter(m instantiate୦୦match୦int,) *instantiate୦୦stream୦int {
	res := make([]int, 0, len(s.eles))
	for i, e := range s.eles {
		if m(e) {
			res = append(res, s.eles[i])
		}
	}

	return instantiate୦୦Of୦int(res)
}

//line stream.go2:46
func (s *instantiate୦୦stream୦int,) Distinct(c instantiate୦୦comparator୦int,) *instantiate୦୦stream୦int {
	res := make([]int, 0, len(s.eles))

	for i := 0; i < len(s.eles); i++ {
		found := false
		for j := i + 1; j < len(s.eles); j++ {
			if c(s.eles[i], s.eles[j]) == 0 {
				found = true
			}
		}
		if !found {
			res = append(res, s.eles[i])
		}
	}
	return instantiate୦୦Of୦int(res)
}

//line stream.go2:64
func (s *instantiate୦୦stream୦int,) Sort(c instantiate୦୦comparator୦int,) *instantiate୦୦stream୦int {
	sort.SliceStable(s.eles, func(i, j int) bool {
		return c(s.eles[i], s.eles[j]) < 0
	})
	return s
}

func (s *instantiate୦୦stream୦int,) Limit(offset int, limit int) *instantiate୦୦stream୦int {
	res := make([]int, 0, len(s.eles))
	for i := range s.eles {
		if i >= offset && len(res) <= limit {
			res = append(res, s.eles[i])
		}
	}
	return instantiate୦୦Of୦int(res)
}

func (s *instantiate୦୦stream୦int,) Skip(num int) *instantiate୦୦stream୦int {
	res := make([]int, 0, len(s.eles))
	for i := range s.eles {
		if i >= num {
			res = append(res, s.eles[i])
		}
	}

	return instantiate୦୦Of୦int(res)
}

func (s *instantiate୦୦stream୦int,) ForEach(f func(e int,)) *instantiate୦୦stream୦int {
	for i := range s.eles {
		f(s.eles[i])
	}
	return s
}

func (s *instantiate୦୦stream୦int,) ToSlice() []int {
	return s.eles
}

func (s *instantiate୦୦stream୦int,) Max(c instantiate୦୦comparator୦int,) int {
	panic("implement me")
}

func (s *instantiate୦୦stream୦int,) Min(c instantiate୦୦comparator୦int,) int {
	panic("implement me")
}

func (s *instantiate୦୦stream୦int,) AnyMatch(m instantiate୦୦match୦int,) bool {
	for _, e := range s.eles {
		if m(e) {
			return true
		}
	}
	return false
}

func (s *instantiate୦୦stream୦int,) AllMatch(m instantiate୦୦match୦int,) bool {
	for _, e := range s.eles {
		if !m(e) {
			return false
		}
	}
	return true
}

func (s *instantiate୦୦stream୦int,) NoneMatch(m instantiate୦୦match୦int,) bool {
	for _, e := range s.eles {
		if m(e) {
			return false
		}
	}
	return true
}

func (s *instantiate୦୦stream୦int,) Count() int {
	return len(s.eles)
}

func (s *instantiate୦୦stream୦int,) FindFirst(m instantiate୦୦match୦int,) int {
	for _, e := range s.eles {
		if m(e) {
			return e
		}
	}
	return s.def
}

func (s *instantiate୦୦stream୦int,) FindLast(m instantiate୦୦match୦int,) int {
	res := make([]int, 0, len(s.eles))
	for _, e := range s.eles {
		if m(e) {
			res = append(res, e)
		}
	}

	if len(res) == 0 {
		return s.def
	}

	return res[len(res)-1]
}

func (s *instantiate୦୦stream୦int,) FindAny(m instantiate୦୦match୦int,) int {
	for _, e := range s.eles {
		if m(e) {
			return e
		}
	}
	return s.def
}

func (s *instantiate୦୦stream୦int,) FindNth(m instantiate୦୦match୦int,) int {
	panic("implement me")
}

func (s *instantiate୦୦stream୦int,) Concat(tail *instantiate୦୦stream୦int,) *instantiate୦୦stream୦int {
	res := make([]int, 0)
	res = append(s.eles, tail.eles...)
	return instantiate୦୦Of୦int(res)
}

func (s *instantiate୦୦stream୦int,) ConcatArray(tail []int,) *instantiate୦୦stream୦int {
	res := make([]int, 0)
	res = append(s.eles, tail...)
	return instantiate୦୦Of୦int(res)
}

//line stream.go2:189
type instantiate୦୦match୦int func(e int,) bool
//line stream.go2:191
type instantiate୦୦comparator୦int func(e1, e2 int,) int
//line stream.go2:193
type _ sort.Float64Slice

//line stream.go2:193
var _ = testing.AllocsPerRun