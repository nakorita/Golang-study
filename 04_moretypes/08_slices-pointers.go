package main

import "fmt"

func main() {
	var c [5]int = [5]int{32, 34, 55}
	var d = [5]int{6, 8, 10}

	e := [...]string{"Maria", "Andrew"}

	fmt.Println(c, d)
	fmt.Println(e)

	e[1] = "John"

	fmt.Println(e)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	p := &names
	fmt.Println(a, b)
	fmt.Printf("확인1 : %p\n", &names)
	fmt.Printf("확인2: %v\n", *p)
	fmt.Printf("확인3 : %v\n", names)
	fmt.Printf("확인4 : %v, %p\n", &a, &a)
	fmt.Printf("확인5 : %v, %p\n", &b, &b)

	// b[0] = "Vivian"
	fmt.Println(a, b)
	fmt.Println(names)

	f := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(f)
	fmt.Println(f[0][1])
	fmt.Println(f[1][1])

}

// 기초 - 더 많은 타입들 - 배열을 참조하는 슬라이스
// 슬라이스는 어떤 데이터도 저장할 수 있다.
// 슬라이스의 요소를 변경하면 기본 배열의 해당 요소도 수정된다.
// 동일한 기본 배열을 공유하는 다른 슬라이스들은
// 변경된 값을 볼 수 있다.

// 중괄호에서 값은 한 줄로 나열해도 되고, 여러줄로 나열해도 된다.
// 여러 줄로 나열할 때는 마지막 요소에도 콤마를 붙여준다.
// ...를 사용하면 초기화할 값의 개수에 따라 자동으로 크기 설정.
