package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	var s1 string = "Press 'Enter' to continue...\v"
	/*
		- a : 경고음, 벨
		- b : 벡스페이스
		- n : 새줄
		- t : 수평탭
		- v : 수직탭
		- \\: 백슬래시를 직접추가
		- \' : '를 문장에 추가
		- \" : 위와 동일
	*/
	var s2 string = `레드벨벳
	영원하라` // 두줄이상시 `` 즉 벡쿼드로 잠그시길
	fmt.Println(len(s1)) // 문자열의 길이
	fmt.Println(len(s2))
	fmt.Println(utf8.RuneCountInString(s2)) // 한글의 실제길이
	fmt.Println(s1 == s2)
	fmt.Println(s1 + s2)
	fmt.Printf("%c\n", s1[3])
	//  s1[0]  =  c : 이런식으로 수정 불가능한 문자열
	fmt.Println(s2 + "  슬기야 내가 많이 좋아해")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
