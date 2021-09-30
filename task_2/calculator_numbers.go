package main

import "fmt"


func main() {
	var a int
	fmt.Println("Введите произвольное число до 1000:")
	fmt.Scan(&a)

	b := a / 100
	c := (a - b * 100) / 10
	d := (a - b * 100 - c * 10) / 1

	fmt.Print( "сотни: ", b, ". десятки: ", c, ". единицы: ", d, ".")

}

