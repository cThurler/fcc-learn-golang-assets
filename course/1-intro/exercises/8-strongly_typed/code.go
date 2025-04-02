package main

import "fmt"

func main() {
	var username string = "wagslane"
	var password int = 20947382822
	var str string = fmt.Sprint(password)
	// correto seria mudar o valor de password para string, mas achei que poderia aprender mais assim.
	// don't edit below this line
	fmt.Println("Authorization: Basic", username+":"+str)
}
