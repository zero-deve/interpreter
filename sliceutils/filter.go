package sliceutils

type FilterFn[I comparable] func(item I) bool

func FilterSlice[I comparable](slice []I, filterFn FilterFn[I]) []I {
	filteredSlice := []I{}

	for _, item := range slice {
		if filterFn(item) {
			filteredSlice = append(filteredSlice, item)
		}
	}

	return filteredSlice
}
