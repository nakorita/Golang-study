package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}

// 기초 - 패키지와 변수,함수 - 함수
// 함수는 0개 혹은 그보다 많이 인자를 받을 수 있다.
// 위의 예시 add()는 2개의 int형 매개변수를 이용.
// ***변수 이름 뒤에 type이 온다는 것을 명심해라.***
