package main

/*     Packages     */
import "fmt"
import "os"
import "bufio"

/*     Main Function     */
func main() {
	//Variables
	scanner := bufio.NewScanner(os.Stdin)

	//Menu
	fmt.Println("Database Management System by Amir Behmaram")
	fmt.Println("Currently will perform: ")
	fmt.Println("	- Database creation & deletion")
	fmt.Println("	- Table creation, deletion, update, and query")
	fmt.Println("To exit, type .q")

	//Polling loop
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
