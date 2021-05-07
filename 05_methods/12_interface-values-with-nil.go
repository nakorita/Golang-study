package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil!>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	fmt.Println("i확인: ", i)
	fmt.Println("&i확인: ", &i)
	describe(i)
	i.M()

	i = &T{"hello"}
	fmt.Println("i확인: ", i)
	describe(i)
	i.M()

}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// 메소드와 인터페이스 - 메소드와 인터페이스 - Nil 인터페이스 값
// 인터페이스 자체 내부의 콘크리트 값이 0일 경우
// 그 메소드는 nil리시버로 호출된다.
// 일부 언어에서는 이게 null포인터 예외를 발생시키지만
// Go에서는 nil리시버로 호출되는 것이고, 이렇게 사용하는 게 일반적.
// ***nil 콘크리트 값을 갖는 인터페이스 값 자체가 nil은 아니다.
