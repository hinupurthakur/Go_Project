package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func Boolean() bool {
	var b bool
	b = true
	return b
}

func showvar() {
	var a string = "ImString"
	fmt.Println(a)
	var e int
	fmt.Println(e) //default value of int is 0

	var b, c int = 1, 2
	fmt.Println(b, c)
	d := "This is Go Lang" //Just to showcase the beauty of := operator
	f := a + " : " + d     //String concatination
	fmt.Println(f)
}
