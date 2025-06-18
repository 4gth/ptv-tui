package joinutil

func MergeToMap[K comparable, V any](
	target map[K]*V,
	source []V,
	keyFunc func(V) K,
	mergeFunc func(dst *V, src V),
) {
	for _, item := range source {
		key := keyFunc(item)
		if existing, found := target[key]; found {
			mergeFunc(existing, item)
		} else {
			copy := item
			target[key] = &copy
		}
	}
}
