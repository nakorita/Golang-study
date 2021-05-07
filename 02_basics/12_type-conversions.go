package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = -3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

// 기초 - 패키지와 변수,함수 - Type 변환
// int와 uint 설명은 기초파트.docx 참고
