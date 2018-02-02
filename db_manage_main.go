package main

/*     Packages     */
import "fmt"
import "os"
import "bufio"

/*     Global Constants    */
var validDbStringsU = []string{"CREATE DATABASE","DROP DATABASE","USE"}
var validDbStringsL = []string{"create database","drop database","use"}

var validTableStringsU = []string {"CREATE TABLE","DROP TABLE","SELECT * FROM","ALTER TABLE"}
var validTableStringsL = []string {"create table","drop table","select * from","alter table"}

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
      line := scanner.Text()
      if line == ".q" {
			fmt.Println("Program ending")
         break
      } else {
			/* Need to break up input line into seperate values:
			 	check first value, if it exists, begin building execution string
					else: Bad input error
				check second value, it it exists, add to exec string
					else: check if it is a current db name (This is assuming first value
							was USE otherwise -> bad input error
				continue checking until ; is hit. If no ; present at end "Missing ; error"
			*/
			//Break input into pieces
			if existsInArr(line[0], validTableStringsL) {
				fmt.Println("Exists in lower case")
			}


		}
   }
}

func breakInput(input string) []string {
	//Temporary array return
	s := make([]string, 3)
	s[0] = "temp"
	s[1] = "temp1"

	return s
}

func validateString(input string) bool {
	return true
}

func existsInArr(testStr string, arr []string) bool{
	for i := 0; i < len(arr); i++ {
		if testStr == arr[i] {
			return true
		}
	}
	return false
}
