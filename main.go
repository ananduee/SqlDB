package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ananduee/SqlDB/compiler"
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
			parsedStatement := compiler.Parse(text)
			switch parsedStatement.Type {
			case compiler.INSERT:
				fmt.Println("This is to handle insert.")
			case compiler.SELECT:
				fmt.Println("This is to handle select.")
			case compiler.UNRECOGNIZED:
				fmt.Printf("SqlDB > Unrecognized command. '%s'\n", text)
			}
		}
	}
}
