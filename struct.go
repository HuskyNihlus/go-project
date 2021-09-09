package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	var flag int
	flag = 1
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	var str = []rune(text)
	for i := 0 ; i < len(str) ; i++ {
		if str[len(str) - (i + 1)] == str[i] && flag == 1 {
			flag = 1
		} else {
			flag = 0
		}
	}
	if flag == 1 {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}