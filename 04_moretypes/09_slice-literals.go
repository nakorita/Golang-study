package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

// 기초 - 더 많은 타입들 - 슬라이스 리터럴
// 슬라이스 리터럴은 길이가 없는 배열 리터럴과 같다.

// 배열 리터럴
// [3]bool{true, true, false}
// 위의 배열을 참조하는 슬라이스
// []bool{true, true, false}
