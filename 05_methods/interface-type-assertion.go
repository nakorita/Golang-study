package main

import "fmt"

func main() {
	var num interface{} = 10

	a := num
	b := num.(int)

	fmt.Printf("%T, %d\n", a, a)
	printTest(b)
}

func printTest(i interface{}) {
	fmt.Printf("%T, %d\n", i, i)

}

// Assertion이란 '주장'이라는 뜻.
// 인터페이스형으로 선언된 변수는
// 초기화한 값에 따라 형이 자동으로 명시된다.
// 즉 dynamic type이다.
// 확실한 형을 표현하기 위해서 'type assertion'을 해준다.
// 인터페이스 형에서 형을 선언하려면
// "변수이름.(형)" 을 명시하면 된다.
// ***인터페이스 형으로 선언됐는데 nil값인 경우 에러 발생.
