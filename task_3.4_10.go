package main

import (
	// 	//"encoding/json" // пакет используется для проверки ответа, не удаляйте его

	// 	"errors"
	"fmt" // пакет используется для проверки ответа, не удаляйте его
	// 	//"os"  // пакет используется для проверки ответа, не удаляйте его
)

// func readTask() (interface{}, interface{}, interface{}) {
// 	return 5.4, 2.0, "+"
// }

// func main() {

// 	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
// 	// все полученные значения имеют тип пустого интерфейса

// 	num1, err := getValue(value1)
// 	if err != nil {
// 		panic(err)
// 	}

// 	num2, err := getValue(value2)
// 	if err != nil {
// 		panic(err)
// 	}

// 	oper, err := getOperation(operation)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("%.4\n", oper(num1, num2))
// }

// func mathOperation(num1, num2 float64, operation string) float64 {

// 	return 0
// }

// func getValue(value interface{}) (float64, error) {
// 	if v, ok := value.(float64); ok {
// 		return v, nil
// 	}
// 	return 0, errors.New(fmt.Sprintf("value=%v: %T", value, value))
// }

// func getOperation(operation interface{}) (fn func(num1, num2 float64) float64, err error) {
// 	v, ok := operation.(string)
// 	if !ok {
// 		return nil, errors.New(fmt.Sprintf("неизвестная операция"))
// 	}

// 	switch v {
// 	case "+":
// 		fn = func(num1 float64, num2 float64) float64 {
// 			return num1 + num2
// 		}
// 	case "-":
// 		fn = func(num1 float64, num2 float64) float64 {
// 			return num1 - num2
// 		}
// 	case "*":
// 		fn = func(num1 float64, num2 float64) float64 {
// 			return num1 * num2
// 		}
// 	case "/":
// 		fn = func(num1 float64, num2 float64) float64 {
// 			return num1 / num2
// 		}
// 	default:
// 		err = errors.New(fmt.Sprintf("неизвестная операция"))
// 	}
// 	return
// }

func readTask() (value1, value2, operation interface{}) {
	return 5.0, 5.6, "/"
}

func printError(value interface{}) {
	fmt.Printf("value=%v: %T\n", value, value)
}

func main() {
	value1, value2, operation := readTask()
	// var v1, v2 float64
	v1, ok := value1.(float64)
	if !ok {
		printError(value1)
		return
	}

	v2, ok := value2.(float64)
	if !ok {
		printError(value2)
		return
	}

	if v, ok := operation.(string); ok {
		switch v {
		case "+":
			result := v1 + v2
			fmt.Printf("%.4f\n", result)
		case "-":
			result := v1 - v2
			fmt.Printf("%.4f\n", result)
		case "*":
			result := v1 * v2
			fmt.Printf("%.4f\n", result)
		case "/":
			result := v1 / v2
			fmt.Printf("%.4f\n", result)
		default:
			fmt.Println("неизвестная операция")
			return
		}
	} else {
		fmt.Println("неизвестная операция")
		return
	}
}
