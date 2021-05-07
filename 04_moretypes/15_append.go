package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// (추가할 슬라이스, 추가할 요소)
	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// 기초 - 더 많은 타입들 - 슬라이스에 요소 추가하기
// Go는 append 함수를 통해 슬라이스에 새로운 요소를 추가.
// ex)
// func append(s []T, vs ...T) []T
// append의 첫 번째 파라미터 s는 슬라이스 타입 T
// 그리고 나머지 T값들을 슬라이스에 추가
// ***append의 결과 값은 원래 슬라이스의 모든 요소와
// 추가로 제공된 값들을 포함하는 슬라이스이다.
// 만약 s의 원래 배열이 너무 작아서 주어진 값을
// 모두 추가할 수 없다면, 더 큰 배열이 할당된다.
// 이 때 반환된 슬라이스는 새로 할당된 배열을 가리킨다.
