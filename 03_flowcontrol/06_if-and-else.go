package main

import (
	"fmt"
	"math"
)

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// return v
	// v 변수는 if문 내에서 선언된 것이므로, 바깥에서 사용 x
	return lim
}

func main() {
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
	fmt.Println("---끝---")
}

// 기초 - 흐름 제어 구문 - If와 else
// 짧은 if문 안에서 선언된 변수들은
// 어떠한 else 블럭에서든 사용이 가능하다.
// pow에 대한 두 가지 호출 모두
// ***main의 시작부분에서 fmt.Println에 대한
// 호출 이전에 그들의 결과를 발표한다.
