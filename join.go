package lo

func Join[J, K, R any](
	left []J,
	right []K,
	match func(J, K) bool,
	mapper func(J, K) R,
) []R {
	var r = make([]R, 0, len(left))

	for _, j := range left {
		for _, k := range right {
			if !match(j, k) {
				continue
			}
			r = append(r, mapper(j, k))
		}
	}

	return r
}
