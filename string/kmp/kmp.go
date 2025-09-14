package main

// 求next前缀表，模式串
func getNext(pattern []byte) []int {
	lps := make([]int, len(pattern))
	index := 0
	for i := 1; i < len(pattern); {
		if pattern[i] == pattern[index] {
			lps[i] = index + 1
			index++
			i++
		} else { //不相等，如果下标不为0，index是前一位下标的前缀表对应的值
			if index != 0 {
				index = lps[index-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

func kmpFunction(pattern []byte, str []byte) int {
	lps := getNext(pattern) //模式串的前缀表
	i, j := 0, 0            //i是主串要找的，j是模式串
	for i < len(str) && j < len(pattern) {
		if str[i] == pattern[j] {
			i++
			j++
		} else {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}

	//for ; i < len(str)&&j<len(pattern); i++ {
	//	if str[i]==pattern[j]{
	//		j++
	//	}else if j!=0{
	//		j=lps[j-1]
	//	}
	//}

	if j == len(pattern) {
		return i - len(pattern)
	}
	return -1
}

//func subString(pattern []string, str []string) int {
//	i, j, k := 0, 0, 0
//	for i < len(str) && j < len(pattern) {
//		if str[i] == pattern[j] {
//			i++
//			j++
//		} else {
//			j = 0
//			k++
//			i = k
//		}
//	}
//	if j == len(pattern) {
//		return i-len(pattern)
//	}
//	return -1
//}

func repeatedSubstringPattern(s string) bool {
	ss := []byte(s)
	if len(ss) == 1 {
		return false
	}
	n := len(ss)
	//求next数组
	next := make([]int, len(ss))
	i := 0
	for j := 1; j < len(ss); {
		if ss[i] == ss[j] {
			next[j] = i + 1
			j++
			i++
		} else {
			if i != 0 {
				i = next[i-1]
			} else {
				next[j] = 0
				j++
			}
		}
	}
	//
	if next[n-1] != 0 && n%(n-next[n-1]) == 0 {
		return true
	}
	return false
}

func main() {
	s := "ababba"
	//getNext([]byte(s))
	repeatedSubstringPattern(s)
}
