package main

import "fmt"

// Функция сложения двух чисел
func Add(a, b int) int {
	return a + b
}

func main() {
	result := Add(5, 4)
	fmt.Println(result)
}
