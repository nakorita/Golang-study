package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Hour(), ":", t.Minute())
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// 기초 - 흐름 제어 구문 - 조건이 없는 Switch
// 조건이 없는 Switch는 switch true와 동일.
// 이 구조는 if-else 체인을 작성하는 데에
// 아주 깔끔한 방식일 수 있다.
