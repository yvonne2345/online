package main

//func reverseWords(s string) string {
//	split := strings.Fields(s)
//	result := make([]string, 0)
//	for i := len(split) - 1; i >= 0; i-- {
//		result = append(result, split[i])
//	}
//	return strings.Join(result, " ")
//}

func main() {
	//str := "  hello world  "
	//str := "the sky is blue "
	//reverseWords(str)

	s := "abcdefg"
	n := 2

	reverseTwo(s, n)
}

func reverseWords(s string) string {
	res := []byte{}

	// 快指针倒序遍历字符串
	slow := len(s) - 1
	for fast := len(s) - 1; fast >= 0; fast-- {
		// 判快指针是否遇到单词尾
		if s[fast] != ' ' && s[slow] == ' ' {
			// 更新慢指针指向单词尾
			slow = fast
		} else if s[fast] == ' ' && s[slow] != ' ' {
			// fast遍历到单词首
			// 当前单词追加到结果集
			res = append(res, []byte(s[fast:slow+1])...)
			// 更新慢指针指向空格
			slow = fast
		}
	}
	// 判是否存在首字符的单词
	if s[0] != ' ' {
		res = append(res, ' ')
		res = append(res, []byte(s[:slow+1])...)
	}
	// 去除首空格后返回
	return string(res[1:])
}

func reverseTwo(s string, n int) string {
	str := []byte(s)
	ans := make([]byte, 0)
	ans = append(ans, str[len(s)-n:]...)
	ans = append(ans, str[:len(s)-n]...)

	return string(ans)
}
