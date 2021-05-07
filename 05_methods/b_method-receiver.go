package main

import "fmt"

type Point2 struct {
	X, Y int
}

// Point2는 두 개의 메소드가 선언된 상태

// pointer receiver
// 이 메소드의 리시버는 pointer 리시버이다.
// 여기서의 복제는 값이 복제되는 게 아니라
// 레퍼런스가 넘어간다.
// 즉 p는 원본 그 자체
// 왜냐하면 포인터로 받았으니까.
func (p *Point2) add(a int) {
	p.X += a
	p.Y += a
}

// value receiver
// 여기서는 값이 복제된다.
// 즉 다른 값이다.
func (p Point2) mul(a int) {
	p.X *= a
	p.Y *= a
}

func main() {
	p := Point2{3, 4}
	p.add(10)
	p.mul(100)
	fmt.Println(p)
}
