package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	v := []int{11, 22, 33}

	if s == nil {
		fmt.Println("It's nil!")
	}
	fmt.Printf("길이: %d, 용량: %d\n", len(v), cap(v))
}

// 기초 - 더 많은 타입들 - nil 슬라이스
// 슬라이스의 제로값은 nil이다.
// nil 슬라이스의 길이와 용량은 0이고,
// 기본 배열을 갖고 있지 않다.
