package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// 메소드와 인터페이스 - 메소드와 인터페이스 - Nil 인터페이스 값
// ***위의 코드 런타임 에러 발생
// nil 인터페이스에서 메소드를 호출하는 것은
// 런타임 에러를 발생시킨다.
// 왜냐하면 어떤 구체적인 메소드를 호출할지를 나타내는
// 인터페이스 튜플 내부의 타입이 없기 때문이다.
