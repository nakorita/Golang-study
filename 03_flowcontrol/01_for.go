package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// 기초 - 흐름 제어 구문 - for
// Go는 for 반복문이라는 단 하나의 반복 구조를 갖는다.
// 기본적인 for 반복문의 ; 으로 구별되는 3가지 구성 요소를 갖는다.
// for문의 구조
// ***초기화 구문; 조건 표현; 사후 구문
// *** C나 Java, JavaScript와 다르게
// Go는 for문의 세 가지 구성 요소를
// 감싸는 괄호가 없고, {} 괄호가 항상 필수입니다.
