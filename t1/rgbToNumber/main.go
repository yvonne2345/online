package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//rgb->num
	tests := []string{"#ffffff", "#FFCCE5", "255, 255, 255"}
	for _, test := range tests {
		//转
		err := convertRgb(test)
		if err != nil {
			errors.New("失败")
		}
	}
}

func convertRgb(str string) error {
	if strings.HasPrefix(str, "#") {
		//16
		err := hexToNumber(str)
		if err != nil {
			return errors.New("失败")
		}
	} else {
		//10
		err := intToNumber(str)
		if err != nil {
			return errors.New("失败")
		}
	}
	return nil
}

func hexToNumber(str string) error {
	hex := []string{}
	num := []int64{}
	str = strings.TrimPrefix(str, "#")
	for i := 0; i < len(str); i = i + 2 {
		hex = append(hex, str[i:i+2])
	}
	//fmt.Println(hex)
	//转int
	for _, value := range hex {
		a, err := strconv.ParseInt(value, 16, 10)
		if err != nil {
			errors.New("失败")
		} else {
			num = append(num, a)
		}
	}
	fmt.Println(num)
	return nil
}

func intToNumber(str string) error {
	num := []int{}
	split := strings.Split(str, ",")
	for _, value := range split {
		atoi, err := strconv.Atoi(strings.TrimSpace(value))
		if err != nil {
			errors.New("失败")
		} else {
			num = append(num, atoi)
		}
	}
	//转int
	fmt.Println(num)
	return nil
}

//package main
//
//import (
//	"errors"
//	"fmt"
//	"strconv"
//	"strings"
//)
//
//// parseRGB 解析 RGB 颜色字符串，支持十六进制和十进制格式
//func parseRGB(input string) (rgb []int64, err error) {
//	input = strings.TrimSpace(input)
//	if strings.HasPrefix(input, "#") {
//		// 处理十六进制格式
//		if len(input) != 7 {
//			return []int64{}, errors.New("invalid hex color length")
//		}
//		hexColor := strings.TrimPrefix(input, "#")
//		for i := 0; i < len(hexColor); i = i + 2 {
//			num, err1 := parseHexComponent(hexColor[i : i+2])
//			if err1 != nil {
//				return []int64{}, err1
//			}
//			rgb = append(rgb, num)
//		}
//		return rgb, nil
//	} else {
//		// 处理十进制格式
//		components := strings.Split(input, ",")
//		if len(components) != 3 {
//			return []int64{}, errors.New("invalid decimal color format")
//		}
//		for _, component := range components {
//			num, err1 := parseDecimalComponent(component)
//			if err1 != nil {
//				return []int64{}, err1
//			}
//			rgb = append(rgb, num)
//		}
//		return rgb, nil
//	}
//}
//
//// parseHexComponent 将两位十六进制字符串转换为整数
//func parseHexComponent(hex string) (int64, error) {
//	value, err := strconv.ParseInt(hex, 16, 32)
//	if err != nil || value < 0 || value > 255 {
//		return 0, errors.New("invalid hex value")
//	}
//	return value, nil
//}
//
//// parseDecimalComponent 将十进制字符串转换为整数
//func parseDecimalComponent(dec string) (int64, error) {
//	value, err := strconv.Atoi(strings.TrimSpace(dec))
//	if err != nil || value < 0 || value > 255 {
//		return 0, errors.New("invalid decimal value")
//	}
//	return int64(value), nil
//}
//
//// 主函数用于测试
//func main() {
//	tests := []string{"#ffffff", "#FFCCE5", "255, 255, 255"}
//	for _, t1 := range tests {
//		rgb, err := parseRGB(t1)
//		if err != nil {
//			fmt.Println("转换失败")
//		} else {
//			fmt.Println(rgb)
//		}
//	}
//
//}
