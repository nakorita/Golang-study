package main

import "fmt"

type person2 struct {
	name    string
	age     int
	contact string
}

func addAgeRef(a *person2) { //Pass by reference
	a.age += 4 //자동 역참조 * 생략
}

func addAgeVal(a person2) { //Pass by value
	a.age += 4
}

func main() {
	var p1 = new(person2) //포인터 구조체 객체 생성
	var p2 = person2{}    // 빈 구조체 객체 생성

	fmt.Println(p1, p2)

	p1.age = 25
	p2.age = 25

	addAgeRef(p1) //&을 쓰지 않아도 됨
	addAgeVal(p2)

	fmt.Println(p1.age)
	fmt.Println(p2.age)

	addAgeRef(&p2) //&을 붙이면 pass by reference 가능
	fmt.Println(p2.age)
}
