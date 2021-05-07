package main

import "fmt"

func getClosure() func() {
	cnt := 0

	return func() { //클로저 함수(익명함수)
		cnt += 1
		fmt.Println(cnt)
	}
}

func main() {
	f := getClosure()
	f()
	f()
	f()
	f()
}
