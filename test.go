package main

import "fmt"

func main() {

	name := "hi" //string
	age := 28 // int
	height := 153.7 //float 64

	kundali := make(map[string]string) // dict
	kundali["rashi"] = "mesh"
	kundali["nakshtra"] = "krithika"

	finance := []int{3000,58000,700} // list

	fmt.Println(kundali)
	fmt.Println(name, age, height)
	fmt.Println(finance)

	for i:=0; i< age; i++ {
		fmt.Println("Nidhi was once: ", i+1 , "year(s) old") // for loop
	}

	for index,value := range finance{
		fmt.Println("index: ", index, "value:", value ) // enumerate kinda
	}

	for key, value := range kundali{
		fmt.Println("key:", key, "value:", value) // iterate over dict
	}

	twoNames := namePrints("Nidhi", "Shriram")
	fmt.Println(twoNames)

	ptr := 899.90
	fmt.Println(&ptr)
	
	x := positiver(&ptr)
	fmt.Println(x)
}

func namePrints(boyName string, girlName string) string{
	return boyName +" x "+ girlName
}

func positiver(num *float64) float64 {
	return *num * *num
}
