package util

func IntContain(i int, set []int) bool {
	for _, i2 := range set {
		if i2 == i {
			return true
		}
	}
	return false
}

func Int64Contain(i int64, set []int64) bool {
	for _, i2 := range set {
		if i2 == i {
			return true
		}
	}
	return false
}
