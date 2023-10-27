package main

import "fmt"

func divide(a int, b int) int {
	return a / b
}

// func main() {
// 	var input int
// 	fmt.Scan(&input)
// 	fmt.Println(divide(input, 5)) //Выведем результат
// }
func main() {
	var input int
	_, err := fmt.Scan(&input) // функция Scan возвращает два параметра, но сейчас важно проверить
	// только ошибку
	if err != nil {
		fmt.Println("Проверьте типы входных параметров")
	} else {
		fmt.Println(divide(input, 5)) //Вывод результата, если ошибок нет
	}
}
