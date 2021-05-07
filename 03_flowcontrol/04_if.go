package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}

// 기초 - 흐름 제어 구문 - If
// Go의 if문은 for 반복문과 비슷하다.
// ()로 둘러쌓일 필요는 없지만,
// {} 괄호는 필수이다.
// fmt.Sprint()는 string타입으로 리턴된다.
// Print나 Fprint의 경우 바이트를 int 타입으로 리턴한다.
