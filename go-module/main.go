package main

import (
	"fmt"
	"soda-meow/surprise"

	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
)

func main() {
	fmt.Println("this is main")
	fmt.Println(sayMeow())
	fmt.Println(sayMeowSurprise())

	say, err := cowsay.Say(
		"Hello",
		cowsay.Type("default"),
		cowsay.BallonWidth(40),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(say)
}

func sayMeowSurprise() string {
	return surprise.Surprise(sayMeow())
}

func SayThisIsBasic() string {
	return "This is basic go"
}
