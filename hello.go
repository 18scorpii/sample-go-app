package main

import "fmt"

func main(){
	fmt.Println("Started Programme with the message !!", PrintMessage())
}

func  PrintMessage() string{
	return "Hello World !!"
}