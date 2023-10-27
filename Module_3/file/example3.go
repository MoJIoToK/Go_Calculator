package main

import "os"

func main() {
	// создаем файл
	os.Create("text.txt")
	// переименовываем файл
	os.Rename("text.txt", "new_text.txt")
	// удаляем файл
	os.Remove("new_text.txt")
	// кстати, os позволяет работать не только с файлами
	// выходим из программы:
	os.Exit(0)
}
