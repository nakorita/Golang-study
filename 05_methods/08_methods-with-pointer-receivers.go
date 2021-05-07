package main

import (
	"fmt"
	"math"
)

type Vertex4 struct {
	X, Y float64
}

func (v *Vertex4) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Vertex4형 변수 v의 제곱근 구하는 메소드
func (v *Vertex4) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex4{3, 4}
	// %+v는 구조체의 필드명까지 포함
	fmt.Printf("Before scalling: %+v, Abs: %v\n", v, v.Abs())

	v.Scale(5)

	fmt.Printf("After scalling: %+v, Abs: %v\n", v, v.Abs())
}

// 메소드와 인터페이스 - 메소드와 인터페이스 - 값 또는 포인터 리시버 선택
// 포인터 리시버를 사용하는 두 가지 이유
// 1) 메서드가 리시버가 가리키는 값을 수정할 수 있다.
// 2) 각각의 메서드 call에서의 value 복사 문제를 피하기 위해서.
// 리시버가 큰 구조체라면 이 방법은 더 효율적이다.
// 메소드는 value receiver 혹은 pointer receiver
// 가 있어야 하지만, 둘 다 혼합되어선 안 된다.
