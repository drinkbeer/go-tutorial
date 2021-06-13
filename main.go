package main

import (
	"fmt"

	"github.com/drinkbeer/go-tutorial/cli"
)

func main() {
	fmt.Println("Entry point of Application")
	c := cli.NewCLI()
	err := c.RunCommand("palindrome-check", "-i=fixture/", "-s=true")
	if err != nil {
		fmt.Println("Error when checking palindrome: ", err.Error())
	}
}
