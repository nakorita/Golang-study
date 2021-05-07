package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	// 2, 3, 4 인덱스만 남기고 1, 2 인덱스 삭제
	// 따라서 len=3, cap=3
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

// 기초 - 더 많은 타입들 - make함수로 슬라이스 만들기
// 슬라이스는 내장된 make 함수로 생성할 수 있다.
// make 함수는 0으로 이루어진 배열을 할당.
// 그리고 해당 배열을 참조하는 슬라이스를 반환.
// 용량을 지정하려면 make함수의 세 번째 인자에 값을 전달
