package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

// 기초 - 흐름 제어 구문 - defer 스택 쌓기
// 연기된 함수 호출들은 스택에 쌓인다.
// 한 함수가 종료될 때 그것의 연기된 함수들은
// 후입선출 순서로 수행된다.
