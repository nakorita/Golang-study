package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "세상"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	// Pi = 5.14
	// fmt.Println(Pi)
}

// 기초 - 패키지와 변수,함수 - 상수
// 상수는 변수처럼 선언되지만 const 키워드와 함께 선언된다.
// 상수는 character 혹은 sring, boolean, 숫자값이 될 수 있다.
// ***상수는 :=를 통해 선언될 수 없다.
