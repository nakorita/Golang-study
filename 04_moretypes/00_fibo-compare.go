package main

import (
	"fmt"
	"time"
)

func fibo1(x int) int {
	if x < 2 {
		return x
	}
	return fibo1(x-1) + fibo1(x-2)
}

func fibo10(x int) int {
	answer := 0
	f0 := 0
	f1 := 1

	if x < 2 {
		return x
	}
	for i := 2; i <= x; i++ {
		// answer값 세팅
		answer = f1 + f0
		// 배열로 생각했을 때 한 칸씩 땡김
		f0 = f1
		f1 = answer
	}
	return answer
}

func main() {
	startTime := time.Now()
	fmt.Println("fibo1: ", fibo1(40))
	elapsedTime := time.Since(startTime)
	fmt.Printf("fibo1 실행시간: %s\n", elapsedTime)

	startTime = time.Now()
	fmt.Println("fibo10: ", fibo10(40))
	elapsedTime = time.Since(startTime)

	fmt.Printf("fibo10 실행시간: %s\n", elapsedTime)
}
