package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	// fmt.Println(i)
	describe(i)
	i.M() // 인터페이스 값으로 메서드 호출

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("%v, %T\n", i, i)

}

// 메소드와 인터페이스 - 메소드와 인터페이스 - 인터페이스의 값
// 인터페이스의 값은 값과 콘크리트 type의 튜플이라고 생각할 수 있다.
// interface는 타입을 저장하는 T와
// 값을 저장하는 V 두가지로 나뉘어 저장된다.
// ex) (vaule, type)
// 인터페이스 값은 특정 기초 콘크리트 유형의 값을 가진다.
// 인터페이스 값으로 메소드를 호출하면
// 기본 형식에 동일한 이름의 메서드가 실행된다.
