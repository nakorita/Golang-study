package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 슬라이스에 제로 값을 준다.
	s = s[:0]
	printSlice(s)

	// 슬라이스의 길이를 늘린다.
	s = s[:4]
	printSlice(s)

	// 밑에 코드 실행시 슬라이스 용량이 6인데,
	// 넘어간다. 에러 발생
	// s = s[:7]
	// printSlice(s)

	// 슬라이스의 첫 번째, 두번째 값을 지운다
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// 기초 - 더 많은 타입들 - 슬라이스의 길이와 용량
// 슬라이스는 length(길이)와 capacity(용량)을 둘 다 가짐.
// 슬라이스의 길이란 슬라이스가 포함하는 요소의 개수
// 슬라이스의 용량이란 슬라이스의 첫 번째 요소부터
// 계산하는 기본 배열의 요소의 개수
// 슬라이스 s의 길이와 용량은 len(s) cap(s)로 얻을 수 있음.
// 슬라이스의 길이를 연장하기 위해서는
// 슬라이스에 충분한 용량이 있다면,
// 다시 슬라이싱 하면 된다.
