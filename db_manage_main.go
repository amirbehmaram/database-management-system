package main

import "fmt"
import "os"
import "bufio"

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Database Management System by Amir Behmaram")
	fmt.Println("Currently will perform: ")
	fmt.Println("	- Database creation & deletion")
	fmt.Println("	- Table creation, deletion, update, and query")
	fmt.Println("To exit, type .q")

		fmt.Print("> ")
		for scanner.Scan() {
			fmt.Print("> ")
      	line := scanner.Text()
      	if line == ".q" {
				fmt.Println("Program ending")
         	break
        }
    }
}
