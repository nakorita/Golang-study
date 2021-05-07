package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// 리시버 인수가 없다.
// Abs함수의 매개변수
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}

// 메소드와 인터페이스 - 메소드와 인터페이스 - method is function
// 위의 코드는 기능 변경없이 일반 함수로 작성된 것.
