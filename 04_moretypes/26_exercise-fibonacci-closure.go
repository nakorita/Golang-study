package main

import "fmt"

func fibonacci() func() int {
	fn0 := 0
	fn1 := 0

	return func() int {
		fn2 := fn1 + fn0
		fn1, fn0 = fn2, fn1
		if fn2 == 0 { // base case
			fn1 = 1
		}
		return fn2
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// 기초 - 더 많은 타입들 - 피보나치 클로저 연습
// 피보나치는 재귀함수
