package main

import "fmt"

func main(){
	for a:='a';a<='z';a++{
		if a%2==0{
			fmt.Printf("%c",a)
		}else{
			fmt.Printf("%c",a -32)
		}
	}
}