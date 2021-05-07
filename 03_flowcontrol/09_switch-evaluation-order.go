package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(time.Saturday)
	fmt.Println("Today is", today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// 기초 - 흐름 제어 구문 - Switch 평가 순서
// switch case는 하나의 케이스라도 성공하면
// 멈추는 식으로 위에서부터 아래로 case를 평가.
// ex)

// switch i {
// case 0:
// case f()	:
// }

// 위의 예시에서 i==0이라면 f는 호출하지 않는다.
