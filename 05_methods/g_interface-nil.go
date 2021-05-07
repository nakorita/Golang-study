package main

import "fmt"

type MyString string

type Rect struct {
	width  float64
	height float64
}

func explain(i interface{}) {
	fmt.Printf("value given to explain function is of type '%T' with value %v\n", i, i)
}

func main() {
	ms := MyString("Hello World!")
	r := Rect{5.5, 4.5}
	explain(ms)
	explain(r)
}

// 빈 인터페이스
// https://hoonyland.medium.com/%EB%B2%88%EC%97%AD-interfaces-in-go-d5ebece9a7ea
// 위의 페이지 참조
