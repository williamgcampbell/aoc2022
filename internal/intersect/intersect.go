package intersect

import "strings"

// StringUnique will find the intersection of all characters in both a and b, and return those characters as a string
func StringUnique(a string, b string) string {
	r := Unique(strings.Split(a, ""), strings.Split(b, ""))
	return strings.Join(r, "")
}

// Unique finds the intersection of all elements in both a and b, returning an array of unique elements.
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
