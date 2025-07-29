package main

import "fmt"

func main(){
	age := 22
	citizen := false

	if age >=18 {
		if citizen{
			fmt.Println("You can vote!")
		}else {
			fmt.Println("You need to be a citizen to vote.")
		}
		} else {
			fmt.Println("Too young to vote")
		}
	}
