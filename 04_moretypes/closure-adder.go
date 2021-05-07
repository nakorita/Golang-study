package main

import "fmt"

func outterSum() func() {
	count := 1

	var innerSum func(k int) int //함수 담을 변수 선언

	innerSum = func(k int) int { //클로저 함수(익명함수)
		if k == 1 {
			return 1
		}
		return k + innerSum(k-1)
	}

	return func() {
		fmt.Println(innerSum(count))
		count += 1
	}
}

func main() {
	f := outterSum()
	//n번 실행했을 때 : 1부터 n까지 합
	f()
	f()
	f()
	f()
	f()
	f()
}
