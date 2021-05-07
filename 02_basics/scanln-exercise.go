package main

import "fmt"

func main() {
	var cd, te int
	var name, addr string
	fmt.Print("번호, 이름: ")
	te, _ = fmt.Scanln(&cd, &name)
	fmt.Println(te)

	fmt.Print("다시 입력-번호, 이름, 주소 ")
	te, _ = fmt.Scanln(&cd, &name, &addr)
	fmt.Println("번호는: ", cd, "이름은: ", name, "주소는: ", addr, "반환 개수는: ", te)
}

// Scanln함수는 입력한 값을 변환하지 못했을 때,
// 변환 성공한 개수를 반환한다.
// ex) te, _ => int형식과 error 형식을 전달한다.

// Scanln은 여러 값을 동시에 입력받을 수 있다.
// 빈칸(스페이스 바)로 값을 구분하고
// 엔터(개행)을 하면 입력이 종료된다.
// 입력받는 변수에 '&'연산자를 붙여 입력받는다.
//	(물론 입력받는 변수는 미리 선언되어야 한다.)
