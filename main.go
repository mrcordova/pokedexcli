package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func main()  {
	scanner := bufio.NewScanner(os.Stdin)
	for  {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		// if !moreCmds {
		// 	break
		// }
		cmd := scanner.Text()
		toLowerCmd := strings.ToLower(cmd)
		words := strings.Fields(toLowerCmd)
		fmt.Println("Your command was:", words[0])

	}
}

