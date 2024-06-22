package main

import (
	"fmt"
	"log"
)

func main() {
	n := ""
	fmt.Println("hotfix")
	fmt.Print("Введите данные: ")
	_, err := fmt.Scanf("%s\n", &n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Вы ввели следующие данные: %s\n", n)
}
