package sliceutils

type CompareSliceFn[I any] func(I, I) bool

func CompareSlice[I any](a []I, b []I, compareFn CompareSliceFn[I]) bool {
	if len(a) != len(b) {
		return false
	}

	for index := 0; index < len(a); index++ {
		if !compareFn(a[index], b[index]) {
			return false
		}
	}

	return true
}
