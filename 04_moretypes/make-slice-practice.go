package main

import "fmt"

func main() {
	s := make([]int, 0, 3) // len=0, cap=3 인 슬라이스 선언
	fmt.Println(s)

	for i := 1; i <= 10; i++ { // 1부터 차례대로 한 요소씩 추가
		s = append(s, i)

		fmt.Println(s)
		fmt.Println(len(s), cap(s)) // 슬라이스 길이와 용량 확인
	}

	fmt.Println(s) // 최종 슬라이스 출력
}

// make() 함수는 "make(슬라이스 타입, 슬라이스 길이, 슬라이스의 용량)" 형태로 선언
// 여기서 용량(Capacity)은 생략해서 선언가능
// 용량 생략시 슬라이스와 똑같은 값으로 선언된다.
// 길이 : 초기화된 슬라이스의 요소 개수.
//		즉, 슬라이스에 5개의 값 초기화 -> 길이 5
//		그 후에 값을 추가하거나 삭제하면 길이 바뀜.
// 용량 : 슬라이스는 배열의 길이가 동적으로 늘어날 수 있어서
// 		길이와 용량을 구분한다.
