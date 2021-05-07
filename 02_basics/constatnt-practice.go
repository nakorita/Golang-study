package main

import "fmt"

const username = "kim"
const nser = 1

func main() {
	var cloud, earp = true, "earp"
	// var grass, seed string = false, "seed"
	const a int = 1
	const b, d = 10, 20 //여러 값을 초기화할 수 있음
	const c = "goorm"

	fmt.Println(cloud, earp)
	// fmt.Println(grass, seed)
	fmt.Println(username)
}

// 상수는 한 번 선언 및 초기화하면 나중에 값을 수정할 수 없다.
// 때문에! 꼭 선언과 동시에 초기화해야 한다. 선언만 한다면 에러 발생
// 초기화 후에 사용하지 않아도 에러가 발생하지 않는다.
// 위에 user 상수 코드 보면 된다.
// 왜 에러가 안 나냐면 상수는 변수와 다르게 명시하는 것 자체가 의미가 있다.
// 그리고 상수는 := 이 용법을 사용할 수 없다.
// 왜냐하면 상수는 const라고 꼭 상수임을 명시해야 하기 때문이다.

// 또한 변수는 동일한 형의 변수는 한 번에 여러개 선언할 수 있다.
// 이때 초기화하는 변수의 개수와 초기화하는 값의 개수가 동일하면 된다.
// ***만약 초기화하지 않는다면 모든 값을 초기화하지 않아야 한다.
