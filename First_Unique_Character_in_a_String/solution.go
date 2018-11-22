package First_Unique_Character_in_a_String


func firstUniqChar(s string) int {
	t := []byte(s)
	strMap := make(map[byte]int)
	for index, value := range t {
		if _, ok := strMap[value]; ok {
			strMap[value] = -1
		} else {
			strMap[value] = index
		}
	}

	ans := len(t) + 1
	for _, v := range strMap {
		if v  >= 0 && ans > v {
			ans = v
		}
	}
	if ans > len(t) {
		return -1
	}
	return ans
}