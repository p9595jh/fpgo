package fpgo

import "sort"

// mapping T to U
func Map[T, U any](convert Pi1[*T, *U]) F {
	return func(s *Shell) *Shell {
		res := []U{}
		for i, item := range s.V.([]T) {
			res = append(res, *convert(i, &item))
		}
		return &Shell{res}
	}
}

// filtering T
func Filter[T any](f Pi1[*T, bool]) F {
	return func(s *Shell) *Shell {
		res := []T{}
		for i, item := range s.V.([]T) {
			if f(i, &item) {
				res = append(res, item)
			}
		}
		return &Shell{res}
	}
}

// slicing T as
func Slice[T any](idxes ...int) F {
	return func(s *Shell) *Shell {
		var a, b int
		var l = len(idxes)
		switch l {
		case 0:
			a = 0
			b = l
		case 1:
			a = idxes[0]
			b = l
		default:
			a = idxes[0]
			b = idxes[1]
		}
		if a < 0 {
			a += l
		}
		if b < 0 {
			b += l
		}
		res := make([]T, b-a)
		copy(res, s.V.([]T)[a:b])
		return &Shell{res}
	}
}

// (newValue, prevValue) -> processedPrevValue
func Reduce[T, U any](reducer Pi2[*T, *U, *U]) F {
	return func(s *Shell) *Shell {
		res := new(U)
		for i, item := range s.V.([]T) {
			res = reducer(i, &item, res)
		}
		return &Shell{res}
	}
}

// sort element a, b
func Sort[T any](sorter P2[*T, *T, bool]) F {
	return func(s *Shell) *Shell {
		ts := s.V.([]T)
		sort.Slice(ts, func(i, j int) bool {
			return sorter(&ts[i], &ts[j])
		})
		return s
	}
}

// deep copying an array
func ArrayCopy[T any]() F {
	return func(s *Shell) *Shell {
		t := s.V.([]T)
		res := make([]T, len(t))
		copy(res, t[:])
		return &Shell{res}
	}
}

// use pure value
func Func[T, U any](f P1[T, U]) F {
	return func(s *Shell) *Shell {
		u := f(s.V.(T))
		return &Shell{u}
	}
}

// true if satisfies at least 1
func Some[T any](f Pi1[*T, bool]) F {
	return func(s *Shell) *Shell {
		for i, item := range s.V.([]T) {
			if f(i, &item) {
				return &Shell{true}
			}
		}
		return &Shell{false}
	}
}

// true if satisfies all
func Every[T any](f Pi1[*T, bool]) F {
	return func(s *Shell) *Shell {
		for i, item := range s.V.([]T) {
			if !f(i, &item) {
				return &Shell{false}
			}
		}
		return &Shell{true}
	}
}

// (index, element, array)
func ForEach[T any](f Pi2V[*T, []T]) F {
	return func(s *Shell) *Shell {
		t := s.V.([]T)
		for i, item := range t {
			f(i, &item, t)
		}
		return s
	}
}

func Reverse[T any]() F {
	return func(s *Shell) *Shell {
		t := s.V.([]T)
		l := len(t)
		res := make([]T, l)
		for i, item := range t {
			res[l-i-1] = item
		}
		return &Shell{res}
	}
}

// []T -> T
func At[T any](i int) F {
	return func(s *Shell) *Shell {
		t := s.V.([]T)
		l := len(t)
		i = (l + i) % l
		return &Shell{t[i]}
	}
}

// append all elements
func Append[T any](elements []T) F {
	return func(s *Shell) *Shell {
		t := s.V.([]T)
		t = append(t, elements...)
		return &Shell{t}
	}
}

// mapping a single element
func MapOne[T, U any](f P1[*T, *U]) F {
	return func(s *Shell) *Shell {
		t := s.V.(T)
		u := f(&t)
		return &Shell{*u}
	}
}
