package main

// 判断 t2 是否是 s 的 字母异位词
func isAnagram(s string, t string) bool {
	var num [26]int
	for _, v := range s {
		num[v-'a']++
	}
	for _, v := range t {
		num[v-'a']--
	}
	if num == [26]int{} {
		return true
	}

	return false
}

func canConstruct(ransomNote string, magazine string) bool {
	var record [26]int
	for _, v := range magazine {
		record[v-'a']++
	}
	for _, v := range ransomNote {
		record[v-'a']--
		if record[v-'a'] < 0 {
			return false
		}
	}

	return true
}

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, str := range strs {
		var num [26]int
		for _, v := range str {
			num[v-'a']++
		}
		m[num] = append(m[num], str)
	}

	ans := make([][]string, 0, len(strs))
	//ans := make([][]string, 0)

	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}

func main() {
	//strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	s := "cbaebabacd"
	p := "abc"
	//isAnagram(s, p)
	//canConstruct(s, p)
	//groupAnagrams(strs)
	findAnagrams(s, p)
}

// 找到 s 中所有 p 的异位词的子串，返回这些子串的起始索引
//func findAnagrams(s string, p string) []int {
//	var num [26]int
//	for _, v := range p {
//		num[v-'a']++
//	}
//	//滑动窗口
//	left := 0
//	var result []int
//	for right, v := range s {
//		num[v-'a']--
//		for num[v-'a'] < 0 {
//			num[s[left]-'a']++ //当num[v-'a'] < 0，说明有s中有p没有的字母
//			left++
//		}
//		if right-left+1 == len(p) {
//			result = append(result, left)
//		}
//	}
//	return result
//}

func findAnagrams(s string, p string) []int {
	var num [26]int
	var ans []int
	for _, v := range p {
		num[v-'a']++
	}

	slow := 0
	for fast, v := range s {
		num[v-'a']--
		for num[v-'a'] < 0 {
			num[s[slow]-'a']++
			slow++
		}
		if fast-slow+1 == len(p) {
			ans = append(ans, slow)
		}
	}
	return ans
}
