package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//value receiver
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// pointer receiver
// 만약 *를 없애면 이 메소드는 더이상 동작 X
// *를 없애면 Scale 함수 내에서는 v의 값을 변경하지만,
// v가 포인터가 아니어서 원래의 값은 변경되지 않는다.

// 하지만 여기서는 포인터를 사용하여
// 변수가 함수에 의해 값이 변경될 수 있게 하였다.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

// 메소드와 인터페이스 - 메소드와 인터페이스 - 포인터 리시버
// 포인터 리시버로 메소드를 선언할 수 있다.
// 포인터 리시버가 있는 메소드는 위의 Scale처럼
// 리시버가 가리키는 값을 수정할 수 있다.
// 메소드는 리시버를 수정해야하므로 pointer receiver가
// value receiver보다 일반적이다.
// ex) 만약 Scale 메서드에서 *를 빼고
// 	value receiver로 선언했다면
// 	value receiver를 사용하면 Scale메서드가
//	원래 Vertex값의 복사본에서 작동한다.
