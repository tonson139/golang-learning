package main

import (
	"fmt"
	"soda-meow/surprise"
)

func main() {
	fmt.Println("this is main")
	fmt.Println(sayMeow())
	fmt.Println(sayMeowSurprise())
}

func sayMeowSurprise() string {
	return surprise.Surprise(sayMeow())
}
