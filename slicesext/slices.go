package slicesext

func Flatten[T any](lists [][]T) (res []T) {
	for _, list := range lists {
		res = append(res, list...)
	}
	return res
}

func Filter[T any](s []T, fn func(T) bool) (res []T) {
	for _, e := range s {
		if fn(e) {
			res = append(res, e)
		}
	}
	return res
}

func Find(s []any, target any) any {
	for _, e := range s {
		if e == target {
			return e
		}
	}
	return nil
}

func RemoveAt[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}

func RemoveAtCopy[T any](s []T, index int) []T {
	s1 := make([]T, 0)
	s1 = append(s1, s[:index]...)
	return append(s1, s[index+1:]...)
}

func ContainsElement[T comparable](s []T, target T) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}

func Intersection[T comparable](s0, s1 []T) []T {
	res := make([]T, 0)
	for _, v := range s0 {
		if ContainsElement[T](s1, v) {
			res = append(res, v)
		}
	}
	return res
}
