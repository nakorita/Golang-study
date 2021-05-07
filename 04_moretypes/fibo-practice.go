package main

import "fmt"

func fibonacci(x int) int { // 피보나치 수열의 k번째 항을 구해주는 함수

	if x < 2 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}

func main() {
	fibo := fibonacci(5)
	fmt.Println(fibo)
}
