package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 写一个函数，把 RGB 颜色字符串，转换成对应的三个整型数字
// 要求：
// - RGB 颜色字符串支持的格式有
//   - 十六进制数字格式，以 "#" 开头，不限制字母大小写，比如 "#FFFFFF"
//   - 十进制数字格式，以 "," 分割，允许逗号前后存在空格，比如 "255,255,255"
//
// - 不合法的颜色字符串需要报错
// - 请尝试提供单元测试代码，尽可能多的考虑边界和异常
// - 请尝试优化代码运行效率，并给出代码运行的时间复杂度分析
// - 允许使用的编程语言：C/C++, Golang, Java
// 示例：
// - "#ffffff" -> (255, 255, 255)
// - "#FFCCE5" -> (255, 204, 229)
// - "255, 255, 255" -> (255, 255, 255)

// parseRGB解析RGB颜色字符串，支持十六进制和十进制格式
func parseRGB(input string) (r, g, b int, err error) {
	input = strings.TrimSpace(input)
	if strings.HasPrefix(input, "#") {
		// 处理十六进制格式
		if len(input) != 7 {
			return 0, 0, 0, fmt.Errorf("invalid hex color length")
		}
		var hexColor string
		hexColor = strings.TrimPrefix(input, "#")
		r, err = parseHexComponent(hexColor[0:2])
		if err != nil {
			return 0, 0, 0, err
		}
		g, err = parseHexComponent(hexColor[2:4])
		if err != nil {
			return 0, 0, 0, err
		}
		b, err = parseHexComponent(hexColor[4:6])
		if err != nil {
			return 0, 0, 0, err
		}
		return r, g, b, nil
	} else {
		// 处理十进制格式
		components := strings.Split(input, ",")
		if len(components) != 3 {
			return 0, 0, 0, fmt.Errorf("invalid decimal color format")
		}
		r, err = parseDecimalComponent(components[0])
		if err != nil {
			return 0, 0, 0, err
		}
		g, err = parseDecimalComponent(components[1])
		if err != nil {
			return 0, 0, 0, err
		}
		b, err = parseDecimalComponent(components[2])
		if err != nil {
			return 0, 0, 0, err
		}
		return r, g, b, nil
	}
}

// parseHexComponent将两位十六进制字符串转换为整数
func parseHexComponent(hex string) (int, error) {
	var value int
	_, err := fmt.Sscanf(hex, "%2x", &value)
	if err != nil {
		return 0, err
	}
	if value < 0 || value > 255 {
		return 0, fmt.Errorf("hex value out of range")
	}
	return value, nil
}

// parseDecimalComponent将十进制字符串转换为整数
func parseDecimalComponent(dec string) (int, error) {
	dec = strings.TrimSpace(dec)
	value, err := strconv.Atoi(dec)
	if err != nil {
		return 0, err
	}
	if value < 0 || value > 255 {
		return 0, fmt.Errorf("decimal value out of range")
	}
	return value, nil
}

// 单元测试
//func main() {
//	//tests := []string{"gp18ar-automated-api-pro-lj", "gp18ar-automated-api-pro", "gp18ar-automated-api"}
//	tests := []string{"gp18ar-automated-api-dev-lj", "gp18ar-automated-api-lj-dev", "gp18ar-automated-api-lj-sit"}
//	//parseRGB(tests)
//	var envList []string
//	for _, v := range tests {
//		if strings.Contains(v, "lj") {
//			after := strings.SplitAfter(v, "-")
//			afterN := strings.SplitAfterN(v, "-", len(after)-1)
//			//处理
//			split := strings.Split(afterN[len(after)-2], "-")
//			if split[0] == "lj" {
//				env := ""
//				env = split[1] + "-lj"
//				envList = append(envList, env)
//			} else {
//				envList = append(envList, afterN[len(after)-2])
//			}
//		} else {
//			after := strings.SplitAfter(v, "-")
//			after1 := strings.Split(v, "-")
//			count := strings.Count(v, "-")
//			envList = append(envList, after[count])
//
//			fmt.Println(after, after1, after[count])
//		}
//	}
//	//for _, container := range tests {
//	//	//1带pro(pro、pro-lj)
//	//	if strings.Contains(container, "pro") {
//	//		index := strings.Index(container, "pro")
//	//		env := container[index:]
//	//		envList = append(envList, env)
//	//
//	//	} else {
//	//		//2 有值
//	//		env := "pro"
//	//		envList = append(envList, env)
//	//	}
//	//}
//
//	fmt.Println(envList)
//	//after := strings.SplitAfter(tests, "-")
//	//after1 := strings.Split(tests, "-")
//	//count := strings.Count(tests, "-")
//	//n := strings.SplitAfterN(tests, "-", count)
//}

func main() {
	ids := []int{1, 2, 3}
	fmt.Println(ids)
}
