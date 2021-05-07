package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	fmt.Println(p)
	fmt.Println(v)
	p.X = 1e9
	fmt.Println(v)
	fmt.Println(p)
	(*p).X = 2e9
	fmt.Println(v)
}

// 기초 - 더 많은 타입들 - pointers to structs
// 구조체 포인터를 통해서 구조체 필드를 접근할 수 있다.
// (*p)로 작성하면, 구조체 포인터 p에서
// 구조체의 x필드에 접근할 수 있다.
// 그러나 Go는 역 참조 명시할 필요 없이
// p.X로 작성할 수 있다.
