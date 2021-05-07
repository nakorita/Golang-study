package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

// 기초 - 패키지와 변수,함수 - 짧은 변수 선언
// ***함수 내에서는 := 라는 짧수 변수 선언은
// 암시적 type으로 var 선언처럼 사용 가능.
// 함수 밖에서는 모든 선언이 키워드
// 함수 안에서 사용하지 않으면 에러 발생
// (var, func, 기타 등등)으로 시작하므로
// := 구문은 사용 불가.
