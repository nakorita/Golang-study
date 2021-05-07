package main

import "fmt"

// 넓이를 가지는 녀석들을 도형이라는 걸로 묶겠다
type Figure interface {
	Area() float64 // 메소드와 반환형 기재
}

type Square struct {
	width  float64
	height float64
}

type Triangle struct {
	width  float64
	height float64
}

// Figure의 인터페이스의 구현
func (s Square) Area() float64 {
	return s.width * s.height
}

func (t Triangle) Area() float64 {
	return t.width * t.height / 2
}

func main() {
	// s,t를 Figure형으로 선언
	// -> 생성자를 호출할 때도 반드시 자신의 타입에 맞춰야 함.
	// 	but, 밑에 s = Square{}타입으로 초기화함.
	var s, t Figure
	s = Square{2, 5}
	t = Triangle{3, 4}

	fmt.Println(s)
	fmt.Println(s.Area())

	fmt.Println(t)
	fmt.Println(t.Area())

	// f를 Figure슬라이스로 선언해서
	// 거기에 Square와 Triangle 넣는 게 가능하다.
	// ****이 이유는 Square와 Triangle 모두
	// Area라는 메소드가 존재하기 때문이다.
	// 이를 interface는 암시적으로 충족된다. 라고 한다.
	// 쉽게 말해 우리가 명시적으로 Square는 Figure에
	// 관련있다. 라고 해주지 않아도 알아서 찾는다.
	f := make([]Figure, 0)
	f = append(f, Square{3, 4})
	f = append(f, Triangle{4, 5})
	for i, value := range f {
		fmt.Println("인덱스번호: ", i, value.Area())
	}
}
