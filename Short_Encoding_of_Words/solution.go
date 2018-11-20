package Short_Encoding_of_Words

import (
	"sort"
	"strings"
)

func minimumLengthEncoding(words []string) int {
	// 先去重并翻转
	tempMap := map[string]int{}
	var nWords sort.StringSlice
	for _, word := range(words){
		l := len(tempMap)
		tempMap[word] = 1
		if len(tempMap) > l {
			c := []rune(word)
			for from, to := 0, len(c)-1; from < to; from, to = from + 1, to - 1{
				c[from], c[to] = c[to], c[from]
			}

			nWords = append(nWords, string(c))
		}
	}

	// 稳定排序
	sort.Stable(nWords)
	nWords = append(nWords, "")
	last := ""
	ans := 0
	for _, v := range(nWords){
		if ! strings.HasPrefix(v, last) {
			ans += len(last)+1
		}
		last = v
	}
	return ans

}



func main() {
	s := []string{"time", "me", "bell"}
	minimumLengthEncoding(s)
}