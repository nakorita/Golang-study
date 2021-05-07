package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

// 메소드와 인터페이스 - 메소드와 인터페이스 - 메소드(continued)
// 구조체가 아닌 형식에 대해서도 메소드를 선언할 수 있다.
// 이 예에서는 Abs 메소드가 있는 숫자 유형 MyFloat 확인 가능.
// 메소드와 동일한 패키지에 유형이 정의된
// 수신자가 있는 메소드만 선언 가능.
// 유형이 다른 패키지(int와 같은 빌트인 유형 포함)
// 에 정의된 리시버로 메소드를 선언할 수 없다.
