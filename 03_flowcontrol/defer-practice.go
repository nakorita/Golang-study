package main

import "fmt"

func main() {
	defer fmt.Println("hello")
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println("j")
		fmt.Println(i)
		defer fmt.Println(i)
	}

	fmt.Println("done")

}

// 우선 defer를 제외한 함수들을 차례대로 실행
// defer구문은 어떤 함수를 리턴하기 직전에
// 실행되므로 for문과 main을 비교하면
// for문 실행하고 리턴후에
// 메인 함수내 나머지 함수를 출력하므로
// main함수 가장 상단에 있는
// "hello"가 가장 마지막에 출력된다.
