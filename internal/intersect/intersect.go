package intersect

import "strings"

func StringUnique(a string, b string) string {
	r := Unique(strings.Split(a, ""), strings.Split(b, ""))
	return strings.Join(r, "")
}

// Unique returns a set of elements that are in both a and b
func Unique[T comparable](a []T, b []T) []T {
	var r []T
	set := make(map[T]struct{})
	for _, v := range a {
		set[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := set[v]; ok {
			r = append(r, v)
			delete(set, v)
		}
	}

	return r
}
