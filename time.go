package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	const (
		obrazec = "2006-01-02 15:04:05"
	)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	firstTime, err := time.Parse(obrazec, s)
	if err != nil {
		panic(err)
	}
	if firstTime.Hour() >= 13 {
		future := firstTime.AddDate(0, 0, 1)
		fmt.Println(future.Format(obrazec))
	} else {
		fmt.Println(firstTime.Format(obrazec))
	}
}

//На стандартный ввод подается строковое представление даты и времени определенного события в следующем формате:
//2020-05-15 08:00:00
//Если время события до обеда (13-00), то ничего не меняется, выводится тоже самая запись.
//Если же событие должно произойти после обеда, то оно переносится на один день вперед в это же время.