package main

import (
	"fmt"
	"retrieval-systems/internal/lab1"
)

func main() {
	var num int
	fmt.Println("Enter lab number")
	fmt.Scanf("%d", &num)
	switch num {
	case 1:
		lab1.Lab1()
	default:
		fmt.Println("None selected")
	}

}
