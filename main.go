package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ananduee/SqlDB/compiler"
	"github.com/ananduee/SqlDB/storage"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	table := storage.NewMemoryTable()
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
				err := table.Insert(parsedStatement.DataRow)
				if err != nil {
					fmt.Println("SqlDB > Failed to insert row. ErroCode = ", err)
				} else {
					fmt.Println("SqlDB > Succesfully inserted row.")
				}
			case compiler.SELECT:
				rows, err := table.GetRows()
				if err != nil {
					fmt.Println("SqlDB > Failed to read rows. ErroCode = ", err)
				} else {
					for _, row := range rows {
						fmt.Printf("(%d, %s, %s)\n", row.ID, row.Username, row.Email)
					}
				}
			case compiler.UNRECOGNIZED:
				fmt.Printf("SqlDB > Unrecognized command. '%s'\n", text)
			}
		}
	}
}
