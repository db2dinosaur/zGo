package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Number of args = ", len(os.Args))
	fmt.Println("Whole thing = ", os.Args)
	fmt.Println("Just the prog = ", os.Args[0])
	fmt.Println("Just the args = ", os.Args[1:])
}
