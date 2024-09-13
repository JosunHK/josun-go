package fputil

type Maybe[a any] struct {
	Value a
	Valid bool
}

func Fmap[a any, b any](f func(a) b, arr []a) []b {
	var result []b
	for _, v := range arr {
		result = append(result, f(v))
	}
	return result
}

func Filter[a any](f func(a) bool, arr []a) []a {
	var result []a
	for _, v := range arr {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func Filter2[a any](f func(a) bool, arr *[]a) []*a {
	var result []*a
	for i, v := range *arr {
		if f(v) {
			result = append(result, &((*arr)[i]))
		}
	}
	return result
}

func Has[a comparable](t a, arr []a) bool {
	for _, v := range arr {
		if v == t {
			return true
		}
	}
	return false
}

func Find[a comparable](f func(a) bool, arr []a) Maybe[a] {
	for _, v := range arr {
		if f(v) {
			return Maybe[a]{
				Value: v,
				Valid: true,
			}
		}
	}
	return Maybe[a]{Valid: false}
}

func IndexOf[a comparable](t a, arr []a) int {
	for i, v := range arr {
		if v == t {
			return i
		}
	}
	return -1
}
