package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("SqlDB > ")
		scanner.Scan()
		text := scanner.Text()
		if text == "exit" {
			fmt.Println("SqlDB > Closing :)")
			break
		} else {
			fmt.Printf("SqlDB > Unrecognized command. '%s'\n", text)
		}
	}
}
