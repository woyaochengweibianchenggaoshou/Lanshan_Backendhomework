package main

import "fmt"

func judge(a int) string {
	if a <= 1 {
		return "不是素数"
	}
	if a == 2 {
		return "是素数"
	}
	if a%2 == 0 {
		return "不是素数"
	}
	for i := 3; i*i <= a; i += 2 {
		if a%i == 0 {
			return "不是素数"
		}
	}
	return "是素数"
}

func main() {
	var a int
	fmt.Println("请输入一个正整数：")
	fmt.Scan(&a)
	s := judge(a)
	fmt.Println(s)
}
