package exercise

func lowToUpCase(in string) string {
	res := make([]byte, 0)

	for _, c := range in {
		if c >= 'a' && c <= 'z' {
			c -= 32
			res = append(res, byte(c))
		}
	}

	return string(res)
}

func mergeIntevals (ints []interval) []interval {
	if ints == nil {
		return nil
	}

	res := make([]interval, 0)

	if len(ints) == 0 {
		return res
	}

	cur := ints[0]

	for i, v := range ints {
		if i == 0 {
			continue
		}
		switch {
		case cur.end < v.st:
			res = append(res, cur)
			cur = v
		case cur.end >= v.st && cur.end < v.end:
			cur.end = v.end
		}
	}

	res = append(res, cur)
	return res
}

func rotateImage(im [][]uint) {
	w := len(im)

	for l := 0; l <= w >> 1; l++ {

		for p := l; p < w - l - 1; p++ {
			// top (l, p) ==> (p, len
			im[p][w-l-1], im[l][p] = im[l][p], im[p][w-l-1]
			// right
			im[w-l-1][w-p-1], im[l][p] = im[l][p], im[w-p-1][w-l-1]

			im [w-p-1][l], im[l][p] = im[l][p], im [w-p-1][l]
		}
	}
}