package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	// 42라는 값과 int형 자료형을 가진
	// 인터페이스형 변수 i
	i = 42
	describe(i)

	// "hello"라는 값과 string 자료형을 가진
	// 인터페이스형 변수 i

	// 따라서 인터페이스는 모든 유형의 값을 가질 수 있다.
	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	// 인터페이스의 value와 type을 출력하는 함수
	fmt.Printf("(%v, %T)\n", i, i)
}

// 메소드와 인터페이스 - 메소드와 인터페이스 - 빈 인터페이스 값
// 0 메서드를 지칭하는 인터페이스 유형을 empty interface라고 한다.
// ex) interface{}
// 빈 인터페이스는 모든 유형의 값을 가질 수 있다.
// (모든 유형은 최소 0개의 메소드를 구현한다.)
// 빈 인터페이스는 알 수 없는 값을 처리하는데 이용된다.
// 예를 들어, fmt.Print는 interface{}타입의
//			어떠한 인수라도 취할 수 있다.
