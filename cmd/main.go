package main

import (
	"fmt"
	"retrieval-systems/internal/lab1"
	"retrieval-systems/internal/lab2"
	"retrieval-systems/internal/lab3"
)

func main() {
	var num int
	fmt.Println("Enter lab number")
	fmt.Scanf("%d", &num)
	switch num {
	case 1:
		lab1.Lab1()
	case 2:
		lab2.Lab2()
	case 3:
		lab3.Lab3()
	default:
		fmt.Println("None selected")
	}

}
