package main

import "fmt"

func main() {
	//生成斐波拉契
	list := []int{0, 1}
	for i := 2; i < 10; i++ {
		list = append(list, list[i-1]+list[i-2])
	}
	fmt.Println(list[1:])

	//查质数
	list2 := []int{}
	for _, num := range list {
		flag := true //默认是质数
		//0、1
		if num < 2 {
			continue
		}
		//其它
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			list2 = append(list2, num)
		}
	}
	fmt.Println(list2)
}

//func main() {
//	n := 1000
//	fib := fibonacci(n)
//	fmt.Printf("第 %d 位的斐波那契数是: %s\n", n, fib.String())
//}
//
//func fibonacci(n int) *big.Int {
//	// 初始化第 0 和第 1 位的斐波那契数
//	if n == 0 {
//		return big.NewInt(0)
//	} else if n == 1 {
//		return big.NewInt(1)
//	}
//
//	a := big.NewInt(0) // F(0)
//	b := big.NewInt(1) // F(1)
//	var c *big.Int
//
//	// 计算从 2 到 n 的斐波那契数
//	for i := 2; i <= n; i++ {
//		c = fibonacci(big.Int).Set(a) // 保存 F(i-2)
//		a.Set(b)                // 更新 F(i-2) 为 F(i-1)
//		b.Set(c.Add(c, b))      // 更新 F(i) = F(i-1) + F(i-2)
//	}
//	return b // 返回 F(n)
//}
