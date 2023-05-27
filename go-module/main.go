package main

import (
	"fmt"
	Surprise "soda-meow/surprise"
)

func main() {
	fmt.Println("this is main")
	fmt.Println(sayMeow())
	fmt.Println(sayMeowSurprise())
}

func sayMeowSurprise() string {
	return Surprise.Surprise(sayMeow())
}
