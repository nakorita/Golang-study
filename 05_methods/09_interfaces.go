package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	fmt.Println(a.Abs())

	a = &v
	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 메소드와 인터페이스 - 메소드와 인터페이스 - interfaces
// 인터페이스 형은 메소드의 시그니처 집합으로 정의된다.
// 인터페이스 유형의 값은 해당 메소드를 구현하는
// 모든 값을 보유할 수 있다.
