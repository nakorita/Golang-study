package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s. \n", os)
	}
}

// 기초 - 흐름 제어 구문 - Switch
// switch구문은 연속적인 if - else
// 구문을 사용하는 짧은 방안이다.
// 값이 조건문과 같은 첫 번째 case를 실행한다.
// Go의 switch는 뒤이어 오는 모든 case를
// 실행하는 것이 아니라 오직 첫 번째로
// 선택된 케이스만 실행한다.
// ***다른 언어들에는 case 마지막에 break 구문 필요.
// 그러나 Go에서는 자동으로 break가 제공된다.
// 또한 Go switch case는 상수일 필요 X, 정수일 필요 X
// runtime 내부에 GOOS라는 상수는(string)
// 프로그램이 돌아가는 OS시스템을 말해준다.
