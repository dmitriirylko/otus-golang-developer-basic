// ДЗ #2 "Базовая работа с Git"
//
// Цель ДЗ - освоение работы с Git репозиторием.
// Программа сохраняет в переменную введенную из консоли строчку и выводит ее
// в консоль в отформатированном виде.
package main

import "fmt"

func main() {
	var name string
	fmt.Print("What is your name? ")
	fmt.Scanln(&name)
	fmt.Printf("Hello, %s!\n", name)
}
