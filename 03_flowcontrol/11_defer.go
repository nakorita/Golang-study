package main

import "fmt"

func main() {

	defer fmt.Println("a")
	fmt.Println("hello")

	defer fmt.Println("b")
	defer fmt.Println("world")
	defer fmt.Println("first")
}

// 기초 - 흐름 제어 구문 - defer
// ***defer문은 자신을 둘러싼 함수가 종료할 때 까지
// 어떠한 함수의 실행을 연기한다.
// 연기된 호출의 인자는 즉시 평가되지만
// 그 함수 호출은 자신을 둘러싼 함수가
// 종료할 때 까지 수행되지 않는다.
