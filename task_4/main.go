package main

import "fmt"

func main() {

	var n int
	fmt.Println("Введите размер массива:")
	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		fmt.Println(fmt.Sprintf("Err: %s", err.Error()))
	}

	slice := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Println("Ведите число: ")
		_, err := fmt.Scanf("%d", &slice[i])
		if err != nil {
			fmt.Println(fmt.Sprintf("Err: %s", err.Error()))
		}
	}
	fmt.Println(fmt.Sprintf("slice %f", slice))

	for i := 1; i < len(slice); i++ {
		current := slice[i]
		j := i
		for j > 0 && slice[j-1] > current {
			slice[j] = slice[j-1]
			j--
		}
		slice[j] = current
	}

	fmt.Printf("Финиш = %f", slice)
}

func readInput() (a float64, b float64, err error) {
	fmt.Println("введите первое значение")
	if _, err = fmt.Scanln(&a); err != nil {
		err = fmt.Errorf("Ошибка считывания первого операнда; %s", err)
		return
	}

	fmt.Println("введите второе значение")
	if _, err = fmt.Scanln(&b); err != nil {
		err = fmt.Errorf("Ошибка считывания второго операнда; %s", err)
		return
	}
	return
}
