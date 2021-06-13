package cli

import (
	"fmt"
)

func main() {
	fmt.Println("Entry point of CLI module")
	c := NewCLI()
	c.RunCommand("palindrome-check", "-i=../fixture/is_parlindrome_strings.txt -s=true")
}
