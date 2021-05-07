package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	//

	var a = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	b := a[1:3]
	fmt.Println(b)

	d := a[1:]
	fmt.Println(d)

	e := a[:2]
	fmt.Println(e)

	var j [10]int
	k := j[0:10]
	fmt.Println(k)
	l := j[:10]
	fmt.Println(l)
	m := j[0:]
	fmt.Println(m)
	n := j[:]
	fmt.Println(n)
	// 위의 슬라이스 표현식 모두 같은 의미.
}

// 기초 - 더 많은 타입들 - 슬라이스 기본값
// 상한 또는 하한을 생략하면, 슬라이싱할 때 기본값 사용 가능.
// 하한 기본 값  : 0, 상한 기본 값 : 슬라이스의 길이
