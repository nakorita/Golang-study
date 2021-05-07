package main

import "fmt"

type Point struct {
	X, Y, Z int
}

// 함수명 앞에 변수명과 타입이 기재되어있음.
// 이를 receiver라고 부른다. (이 리시버는 value receiver)
// 이 메소드는 Point의 식구가 된다.
// 식구가 된다는 말은 private 변수에 마음껏 접근 가능하다.
// 원래 다른 함수에서 private 변수에는 접근할 수 없다.
// 이는 Point 타입의 리시버를 달고 있으므로 Point 타입 메소드
func (p Point) printInfo() {
	fmt.Println(p.X)
	fmt.Println(p.Y)
	fmt.Println(p.Z)
}

func main() {
	p := Point{3, 4, 5}
	p.printInfo()
}
