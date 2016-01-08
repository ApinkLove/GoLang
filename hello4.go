package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const (
		Sunday = iota * 30
		/*
			첫번째것만 해주면 자동으로 매겨주고요.
			 0부터 시작하는데
			저런식으로 하면 0 30 이렇게 되겠네요.
		*/
		_ // iota시에 순서번호는 남기면서 내용은 생략시키는 경우입니다.
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		numberOfDays
	)
	fmt.Println(Thursday)
	fmt.Println("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
