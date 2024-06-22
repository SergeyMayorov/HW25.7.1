package main

import (
	"fmt"
	"log"
)

func main() {
	n := ""
	fmt.Println("develop")
	fmt.Print("Введите целое число: ")
	_, err := fmt.Scanf("%s\n", &n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели число: %s\n", n)
}
