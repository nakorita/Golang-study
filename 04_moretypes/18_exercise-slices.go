package main

import (
	"fmt"
)

// Pic이라는 함수를 구현해라.
// 1) 이 함수는 dy개 만큼의
// 길이를 가지는 슬라이스를 리턴해야 한다.
// 2) 또한 슬라이스 각각의 요소들은
// 부호없는 정수 타입 8비트 크기 dx개를 가진다.

func Pic(dx, dy int) [][]uint8 {
	// dy개 만큼의 길이를 가지는 슬라이스를 만든다.
	var arr = make([][]uint8, dy)

	for x := range arr {
		// 각 요소들은 dx개의 8비트 부호없는
		// 8비트 정수 타입을 가지는 슬라이스여야 한다.

		// 즉 dy 슬라이스 먼저 생성,
		// 각 요소마다 dx의 슬라이스를 생성.
		arr[x] = make([]uint8, dx)
		for y := range arr[x] {
			arr[x][y] = uint8(x ^ y)
		}
	}

	return arr
}

func main() {
	fmt.Println(Pic(3, 4))
}

// 기초 - 더 많은 타입들 - 슬라이스 연습하기
