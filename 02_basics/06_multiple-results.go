package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

// 상무님이 설명해주신 코드가 하단 소스코드에 해당됨.
// swap()와 왜 복수개의 결과를 리턴하는지 설명들음.
// func swap(x, y *string) (*string, *string) {
// 	temp := x
// 	x = y
// 	y = temp
// 	return y, x
// }

// func main() {
// 	// a, b := swap("hello", "world")
// 	a := "hello"
// 	b := "world"
// 	fmt.Println(a, b)
// 	swap(&a, &b)
// 	fmt.Println(a, b)
// }

// 기초 - 패키지와 변수,함수 - 복수개의 결과
// 한 함수는 몇 개의 결과든 반환할 수 있다.
//  swap 함수는 두 개의 string을 반환한다.
