package basic

import (
	"fmt"
)

func forloop() string {
	fmt.Println("Try some basic Loop syntax")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	return "Go has only for - no while"
}
