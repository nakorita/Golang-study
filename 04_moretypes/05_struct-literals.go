package main

import "fmt"

type Vertex struct {
	X, Y int
}

type dog struct {
	name  string
	age   int
	color string
}

func addAge(a dog) {
	a.age += 10
}

func minusAge(a *dog) {
	a.age -= 10
}

var (
	v1    = Vertex{1, 2}
	v2    = Vertex{X: 1}
	v3    = Vertex{}
	p     = &Vertex{1, 2}
	v4    = Vertex{}
	p2    = &Vertex{Y: 3}
	p3    = p2
	p4    = &Vertex{}
	tores = dog{}
)

func main() {
	fmt.Println(v1, p, v2, v3, v4, p2, p3, p4)
	bongbong := dog{color: "brown", name: "bongbong"}
	fmt.Println(bongbong)
	fmt.Println(tores)
	var kong = dog{"kong", 10, "black"}
	fmt.Println(kong)
	kong.color = "white"
	fmt.Println(kong)

	nicky := dog{name: "nicky", color: "gray"}
	addAge(nicky)
	fmt.Println(nicky)
	minusAge(&nicky)
	fmt.Println(nicky)
}

// 기초 - 더 많은 타입들 - Structs Literals
// 구조체 리터럴은 필드 값을 나열하여
// 새로 할당된 구조체 값을 나타낸다.
// Name: 구문으로 필드의 하위 집합만 나열 가능.
// (명명된 필수의 순서는 무관하다.)
// 접두사 &은 구조체 값으로 포인터를 반환한다.
