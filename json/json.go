package main

import (
	"fmt"
	"reflect"
)

type J struct {
	a string //小写无tag
	b string `json:"B"` //小写+tag
	C string //大写无tag
	D string `json:"DD" otherTag:"good"` //大写+tag
}

func printTag(stru interface{}) {
	t := reflect.TypeOf(stru).Elem()
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("结构体内第%v个字段 %v 对应的json tag是 %v , 还有otherTag？ = %v \n", i+1, t.Field(i).Name, t.Field(i).Tag.Get("json"), t.Field(i).Tag.Get("otherTag"))
	}
}

func main() {
	//j := J{
	//	a: "1",
	//	b: "2",
	//	C: "3",
	//	D: "4",
	//}
	//printTag(&j)

	ch := make(chan int) // ch 逃逸到堆（被多 goroutine 共享）
	go func() {
		ch <- 1 // 发送的值是字面量，直接复制到通道，不会逃逸
		ch <- 2
		ch <- 3
		close(ch)
	}()
	for {
		n, ok := <-ch // n 是局部变量，但传递给 fmt.Println 时逃逸
		if !ok {
			break
		}
		fmt.Println(n) // n 逃逸到堆（接口类型导致装箱）
	}
	fmt.Println("done")

}
