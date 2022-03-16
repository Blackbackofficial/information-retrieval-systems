package main

import (
	"fmt"
	"retrieval-systems/internal"
)

func main() {
	var num int
	fmt.Println("Enter lab number")
	fmt.Scanf("%d", &num)
	switch num {
	case 1:
		internal.Lab1()
	default:
		fmt.Println("None selected")
	}

}
