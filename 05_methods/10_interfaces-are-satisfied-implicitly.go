package main

import "fmt"

// interface I형 선언
type I interface {
	M()
}

// T형 구조체 선언
type T struct {
	S string
}

// T형 구조체가 갖고있는 메소드 M()
// T형이 인터페이스 I를 구현함을 의미.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	// interface형 변수 i 선언
	// i에 구조체 T형 초기화
	var i I = T{"hello"}
	// interface i에 있는 메소드 M 호출
	i.M()
}

// 메소드와 인터페이스 - 메소드와 인터페이스 - 인터페이스의 암시적 구현
// 인터페이스는 메소드를 실행함으로써 구현한다.
// 인터페이스는 내용을 따로 선언하지 않아도 형으로서 사용 가능.
