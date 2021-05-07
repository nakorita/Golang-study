package main

import "fmt"

type Vertex struct {
	X, Y float64
}

// 리시버 인수가 있는 함수 즉 메소드
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 리시버 인수가 없는 함수 즉 메소드가 아님
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	// v라는 변수에 담긴 Vertex형 값
	v := Vertex{3, 4}
	fmt.Println(v)
	v.Scale(2)
	fmt.Println(v)
	ScaleFunc(&v, 10)
	fmt.Println(v, "\n")

	// p라는 변수에 Vertex{4,3}의 주소를 가리키는 포인터
	p := &Vertex{4, 3}
	fmt.Println(p)
	p.Scale(3)
	fmt.Println(p)
	ScaleFunc(p, 8)
	fmt.Println(p, "\n")

	fmt.Println(v, p)
}

// 메소드와 인터페이스 - 메소드와 인터페이스 - 메소드와 포인터 indirection
// 포인터 리시버가 있는 메소드는
// ***호출될 때 값이나 포인터를 리시버로 받아들인다.
// 위에서 v.Scale(2)라는 문장의 경우
// v가 포인터가 아니라 값인데도
// 포인터 리시버가 있는 메소드가 자동으로 호출됨.
// 즉, Scale 메소드가 포인터 리시버를 가졌기 때문에
// 편의상 Go는 v.Scale(5)라는 것을
// => (&v).Scale(5)로 해석한다.
