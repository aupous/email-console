package array

func SafeIndexAt(arr []string, ind int) string {
	if ind < 0 {
		return ""
	}
	if len(arr) <= ind {
		return ""
	}
	return arr[ind]
}

func StringArrayEqual(first, second []string) bool {
	if len(first) != len(second) {
		return false
	}
	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}
