package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
var name = []string{"apeach", "lion", "muzi", "con"}

func main() {
	fmt.Printf("pow 길이 확인: %v\n 용량 확인: %v\n", len(pow), cap(pow))

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	for j, k := range name {
		fmt.Printf("인덱스: %d 이름: %s\n", j, k)
	}

	whale := pow[:6]
	fmt.Printf("슬라이스: %d\n", whale)

}

// 기초 - 더 많은 타입들 - Range
// for에서 range는 슬라이스 맵 or 맵의 요소를 순회
// 슬라이스에서 range를 사용하면,
// ***각 순회마다 두 개의 값이 반환된다.
// ***첫 번째는 인덱스, 두 번째는 해당 인덱스의 복사본
