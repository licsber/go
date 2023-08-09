package lArray

func StringArrayContains(arr []string, cmp string) bool {
	for _, v := range arr {
		if v == cmp {
			return true
		}
	}

	return false
}
