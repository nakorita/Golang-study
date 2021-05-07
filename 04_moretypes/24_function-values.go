package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

// 기초 - 더 많은 타입들 - 함수 값
// 함수들도 값이다. 다른 값들과 마찬가지로 전달될 수 있다.
// 함수 값은 함수의 인수나 반환 값으로 사용될 수 있다.
