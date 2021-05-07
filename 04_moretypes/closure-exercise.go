package main

import "fmt"

func voltage() func(i int, r int) int {
	return func(i int, r int) int { // 클로저 함수
		return i * r
	}
}

func main() {
	fmt.Println("클로저")
	// 매개변수 a,b를 더한 값을 반환하는 함수 값을 변수에 담음.
	sum := func(a int, b int) int {
		return a + b
	}
	// sum함수에 1,2를 인자로 전달하고 리턴받은 값을 담는 변수 선언
	result := sum(1, 2)
	fmt.Println(result)

	fmt.Println(("-----------"))
	// voltage함수를 담는 cal 변수 선언
	cal := voltage()
	// cal함수에 2,4를 인자로 전달하고 리턴받는 값을 담는 변수 선언
	fmt.Println(cal(2, 4))
}
