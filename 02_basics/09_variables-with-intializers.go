package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

// 기초 - 패키지와 변수,함수 - 변수 초기화
// 변수 선언은 한 변수 당 하나의 초기값 포함 가능.
// 만약 초기값이 존재한다면, type은 생략될 수 있다.
// 이 경우, 변수는 초기값의 type을 취한다.
