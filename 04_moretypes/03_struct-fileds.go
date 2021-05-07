package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

// 기초 - 더 많은 타입들 - Structs Fields
// 구조체의 필드는 .(dot)으로 접근 가능.
