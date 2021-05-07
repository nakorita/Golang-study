package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

// 기초 - 패키지와 변수,함수 - 이어서 Function
// 두 개 이상의 연속된 이름이
// 주어진 함수의 매개변수가 같은 type일 때
// 마지막 변수를 제외한 매개변수들의
// type을 생략할 수 있다.
// x int, y int를 x, y int로 줄였다.
