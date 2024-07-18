package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println(&ptr )

	meraNaam := "Lola Laila"

	var newPtr = &meraNaam

	fmt.Println(newPtr)
	fmt.Println(*newPtr + "_nono")
}