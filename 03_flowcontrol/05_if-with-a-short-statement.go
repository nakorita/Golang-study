package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	// return v
	// 위에처럼 v를 리턴하면 정의되지 않았다고 뜸.
	// 왜냐하면 v는 if문내에서만 존재하는 지역변수
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

// 기초 - 흐름 제어 구문 - 짧은 구문의 If
// math.Pow(x, y)에서 x는 밑 y는 지수
