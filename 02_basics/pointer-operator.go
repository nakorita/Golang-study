package main

import "fmt"

func main() {
	var num int = 5
	var pnum = &num

	fmt.Println("num : ", num)   //num 값
	fmt.Println("pnum :", pnum)  //num의 메모리 주소
	fmt.Println("pnum :", *pnum) //num의 주소로 메모리에 할당돼있는 값 접근

	*pnum++
	fmt.Println("num : ", num)
	fmt.Println("pnum :", *pnum)
	//포인터 연산자를 이용한 값 변경
}
