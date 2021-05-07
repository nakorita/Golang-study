package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		fmt.Println(sum)
		sum += sum
	}
	fmt.Println(sum)
}

// 기초 - 흐름 제어 구문 - 이어서 For
// 초기화 구문과 사후 구문은 필수가 아니다.
// ;을 생략할 수 있다는 점에서 C 의 while은
// Go에서 for로 쓰인다. (for is hos while)
// cf) for는 반복횟수를 명확히 알고 있을 때
//		while은 조건에 따라 반복횟수를 결정해야할 때
//		while, do-while은 조건을 먼저 검사하느냐 차이
