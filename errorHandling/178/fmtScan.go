package main

import "fmt"

func main() {
	var answer1, answer2, answer3 string

	fmt.Print("Name: ")
	//fmt.scan() opens up the terminal
	// allows for user interaction without BASH
	// omg omg omg omg
	_, err := fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}
	fmt.Print("Fav food: ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}
	fmt.Print("Fav Sport: ")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		panic(err)
	}
	fmt.Println(answer1, answer2, answer3)
}
