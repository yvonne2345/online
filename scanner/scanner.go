package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 用go实现第一行输入为整数数组 Array 的长度 n，接下来 n 行，每行一个整数，表示数组的元素。随后的输入为需要计算总和的区间下标：a，b （b > = a），直至文件结束。
func main1() {
	scanner := bufio.NewScanner(os.Stdin)

	// 读取数组长度
	scanner.Scan()
	lenArr, _ := strconv.Atoi(scanner.Text())

	// 读取数组元素
	array := make([]int, lenArr)
	for i := 0; i < lenArr; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		array[i] = num
	}

	//下标
	for scanner.Scan() {
		line := scanner.Text()
		num := strings.Split(line, " ")
		sum := 0
		if len(num) != 2 {
			fmt.Println("Invalid range input")
			continue
		}

		for _, v := range num {
			a, _ := strconv.Atoi(v)
			b, _ := strconv.Atoi(v)

			if a < 0 || b >= lenArr || a > b {
				fmt.Println("Invalid range:", a, b)
				continue
			}
			//计算
			for i := a; i <= b; i++ {
				sum += array[i]
			}
			fmt.Println(sum)
		}
	}
}

// 从键盘输入
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	//获取m、n
	scanner.Scan()
	content := scanner.Text()
	input := strings.Split(content, " ")
	fmt.Println(input)

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")
		fmt.Println(nums)
	}
}
