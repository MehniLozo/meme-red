package tree

type Analog func(x, y interface{}) int

func StringCompare(a, b interface{}) int {
	s1 := a.(string)
	s2 := b.(string)
	min := len(s2)

	if len(s1) < len(s2) {
		min = len(s1)
	}

	diff := 0

	for x := 0; x < min && diff == 0; x++ {
		diff = int(s1[x]) - int(s2[x])
	}

	if diff == 0 {
		diff = len(s1) - len(s2)
	}
	if diff < 0 {
		return -1
	}
	if diff > 0 {
		return 1
	}
	return 0
}

func CompareInt(a, b interface{}) int {
	a1 := a.(int)
	b1 := b.(int)
	switch {
	case a1 > b1:
		return 1
	case a1 < b1:
		return -1
	default:
		return 0
	}
}
