package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}

// 기초 - 패키지와 변수,함수 - 이름이 주어진(기명의) 반환 값
// Go의 반환 값은 이름이 정해질 수도 있다.
// 그 경우, 함수의 맨 위에서 정의된 변수처럼 다뤄진다.
// 이러한 이름들은 반환 값의 의미를 설명하는 데 사용되어야 한다.
// ***인자가 없는 return문은 이름이 주어진 반환값을 반환.
// 이것을 "naked" return이라고 합니다.
// 위의 예제처럼 naked return문은
// 짧은 함수에서만 사용되어야 한다.
// 긴 함수에서는 그것이 가독성을 떨어뜨릴 수 있다.
