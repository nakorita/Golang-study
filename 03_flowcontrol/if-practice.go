package main

import "fmt"

func calc() {
	var opt int
	var num1, num2, result float32

	fmt.Print("1.덧셈 2. 뺄셈 3. 곱셈 4. 나눗셈 선택: ")
	fmt.Scan(&opt)
	fmt.Print(num1, num2)
	fmt.Print("\n두 개의 실수 입력: ")
	fmt.Scan(&num1, &num2)

	if opt == 1 {
		result = num1 + num2
	} else if opt == 2 {
		result = num1 - num2
	} else if opt == 3 {
		result = num1 * num2
	} else if opt == 4 {
		result = num1 / num2
	}

	fmt.Printf("결과: %f\n", result)
}

func main() {
	var num int

	fmt.Print("정수입력 :")
	fmt.Scan(&num)

	if val := num * 2; val == 2 {
		fmt.Print("hello\n")
	} else if val := num * 3; val == 6 {
		fmt.Print("world\n")
	} else {
		fmt.Print("worng number..\n")
	}
	// fmt.Println(val)
	calc()
}

// 조건식 전에 정의된 변수는
// 해당 조건문 블록에서만 사용할 수 있어서
// main 제일 밑줄에 주석 처리한 코드에서 에러난다.
// (정의되지 않은 변수)
