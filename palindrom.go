package main

import (
	"fmt"
)

func main() {
	var (
		text string
		flag bool
	)
	fmt.Scan(&text)
	tr := []rune(text)
	for i, j := len(tr) - 1, 0; i > len(tr) / 2; i-- {
		if tr[i] == tr[j] {
			flag = true
		} else {
			flag = false
			fmt.Print("Нет")
			break
		}
		j++
	}
	if flag {
		fmt.Print("Палиндром")
	}
}