package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// value receiver
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 메소드가 아닌 함수
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}

// 메소드와 인터페이스 - 메소드와 인터페이스 - 메소드와 포인터 indirection2
// 앞의 예제와 동등한 일은 역방향에서 일어날 수 있다.
// 값 인수를 사용하는 함수는 특정 유형의 값을 사용해야 한다.
// ex) var v Vertex
//		fmt.Println(AbsFunc(v)) <- O
//		fmt.Println(AbsFunc(&v)) <- Complile error

// value receiver가 있는 메소드는 다음과 같이 호출될 때,
// 값이나 포인터를 리시버로 사용한다.
// var v Vertex
// fmt.Println(v.Abs()) // OK
// p := &v
// fmt.Println(p.Abs()) // OK
// 이 경우에, p.Abs() 라는 메서드는 (*p).Abs() 로 해석됩니다.
