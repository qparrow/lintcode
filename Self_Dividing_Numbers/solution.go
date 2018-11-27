package Self_Dividing_Numbers


import (
	"strconv"
)

func selfDividingNumbers(left int, right int) []int {
	returnData := []int{}
	var needAdd bool
	for k := left; k <= right; k++ {
		if k % 10 == 0 {
			continue
		}
		needAdd = true
		for _, i := range (strconv.Itoa(k)) {
			t, _ := strconv.Atoi(string(i))
			if ( t == 0) || (k % t != 0) {
				needAdd = false
				break
			}
		}
		if needAdd {
			returnData = append(returnData, k)
		}
	}
	return returnData
}