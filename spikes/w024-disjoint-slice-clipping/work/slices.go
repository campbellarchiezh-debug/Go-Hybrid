package slices

// SplitDisjoint splits slice s at index mid into two provably disjoint slices
// using 3-index capacity clipping: s[0:mid:mid] and s[mid:len(s):cap(s)].
func SplitDisjoint[T any](s []T, mid int) ([]T, []T) {
	if mid < 0 || mid > len(s) {
		panic("slice: split index out of bounds")
	}
	left := s[0:mid:mid]
	right := s[mid:len(s):cap(s)]
	return left, right
}
