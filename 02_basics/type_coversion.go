package main

func main() {
	var i1 int = 100
	var u1 uint = uint(i1)
	var f1 float32 = float32(i1)
	println(f1, u1)

	str1 := "ABC"
	bytes := []byte(str1)
	str2 := string(bytes)
	println(bytes, str2)
}

// 코드 참고 블로그
// http://golang.site/go/article/5-Go-%EB%8D%B0%EC%9D%B4%ED%83%80-%ED%83%80%EC%9E%85
