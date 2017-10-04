package hashtablitza

// HashIndex returns the index from a 5 string array
func HashIndex(arr [5]string, key string) int {
	var returnVal int
	for i := (len(arr) / 2); i > 0; i-- {
		if arr[i] == key {
			returnVal = i
		} else {
			loc := len(arr) / 2
			if arr[loc+i] == key {
				returnVal = loc + i
			} else if arr[loc-i] == key {
				returnVal = loc - i
			}
		}
	}
	return returnVal
}
