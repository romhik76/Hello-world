package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	UserName = "Roman"
)

func getGreetingMessage(names []string) (string, error) {

	if len(names) == 0 {
		return "", fmt.Errorf("Нет больше имен")
	}

	for _, name := range names {
		if name == "" {
			return "", fmt.Errorf("Строка имени пуста")
		}
	}

	s := fmt.Sprintln("hallo " + strings.Join(names, ", ") + "!")
	s += fmt.Sprintln("task_2")
	return s, nil
}

func SayHello() {
	var message string
	var err error

	if message, err = getGreetingMessage([]string{UserName, "Alex", "Sergey", "Artemis", "Evgen"}); err != nil {
		log.Printf(err.Error())
		os.Exit(1)
	}
	fmt.Println(message)

}

func main() {
	SayHello()

	fmt.Println("Введите два операнта (через пробел)")
	var operand1, operand2 float64
	_, err := fmt.Scanf("%f %f", &operand1, &operand2)
	if err != nil {
		fmt.Print("введен неверый операнд")
		return
	} // todo validate

	fmt.Println("Введиет операцю:")
	var operation string
	_, err = fmt.Scan(&operation)
	if err != nil {
		fmt.Print("Операция не введена")
		return
	} // todo validate

	switch operation {

	}
	var result float64
	if operation == "+" {
		result = operand1 + operand2
	} else if operation == "-" {
		result = operand1 - operand2
	} else if operation == "*" {
		result = operand1 * operand2
	} else if operation == "/" {
		result = operand1 / operand2
		if operand1 == 0 || operand2 == 0 {
			fmt.Print("Даная операция с цифрой 0 невозможна")
			return
		} //  todo validate
	} else {
		fmt.Println("Введена неизвестная операция")
		os.Exit(1)
	}

	fmt.Printf("Результат = %5.2f\n", result)
}
