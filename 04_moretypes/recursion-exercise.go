package main

import "fmt"

func fact(n int) int {
	if n == 0 { // base case
		return 1
	}

	if n == 1 {
		return 1
	}

	return n * fact(n-1)
}

func main() {
	fmt.Println("재귀 함수")
	// fact()에 인자로 5 전달
	fmt.Println(fact(5))

	for i := 0; i <= 5; i++ {
		cnt := fact(i)
		fmt.Println("테스트 출력: ", i, "회 ", cnt)
	}
}
