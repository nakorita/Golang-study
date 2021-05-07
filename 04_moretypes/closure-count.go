package main

func hitCount() func() int { // hitCount()를 몇번 호출해는지 알려주는 함수
	// 지역변수 cnt 선언 및 초기화
	cnt := 0
	// int형 정수를 리턴하는 익명함수
	var addOne func()
	addOne = func() {
		cnt += 1
	}

	return func() int { // 클로저 함수 cnt를 1증가하고 반환하는 함수

		addOne() // cnt에 1을 더한다

		return cnt
	}
}

// 1) 변수 관리를 하려고 일단 전체를 한 함수로 감싸줌
// func hitCount()  func() int{
// 2) 이 함수내에서 관리하는 지역변수 선언
// 	count:=0

// 얘가 hitCount() 의 핵심 기능
// 3) 내가 맨 처음에 의도했던 기능이 있는 함수
// 	return func() int { // 클로저 함수 cnt를 1증가하고 반환하는 함수
// 		count+=1
// 	}
// }

func hitCount() func() int{
	count := 0
	// function이라는 함수형 변수를 선언함.
	var function func() int
	function = func() int{
		return 1
	}
	return function
}

// 밑의 f는 변수가 아니라 함수라고 선언한 것.
// a랑 아예 다름.
func f() {

}

// a라는 변수에 함수를 담을 수 있는데, 이건 곧 go언어가 
// 1급함수라는 개념이 있어서이다.
// 1급함수는 함수를 값으로 취급할 수 있다는 의미.
a := func(){

}


func main() {
	f()
	a()
}
