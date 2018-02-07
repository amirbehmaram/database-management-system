package main

/*     Packages     */
import (
	"fmt"
	"os"
	"bufio"
   "strings"
)

/*     Global Constants    */
var validDbStringsU = []string{"CREATEDATABASE","DROPDATABASE","USE"}
var validDbStringsL = []string{"createdatabase", "dropdatabase","use"}

var validTableStringsU = []string {"CREATETABLE", "DROPTABLE", "SELECT*FROM","ALTERTABLE"}
var validTableStringsL = []string {"createtable", "droptable","select*from","altertable"}

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
			// If input length is greater than 0, start query, else output >
			inputLength := len(line)
			if inputLength > 0 {
				// make sure last value is ;
				if string(line[inputLength - 1]) == ";" {

					// split up text string
					text := strings.Split(line, " ")

					// get table/db name, will have ; at end
					varName := text[len(text) - 1]

					// build query string
					var queryString string
					for i := 0; i < len(text); i++ {
						tempStr := text[i]

						if tempStr != varName {
							queryString += tempStr
						} else {
							break
						}
					}

					// hack ; off the end of varName
					if last := len(varName) - 1; last >= 0 && varName[last] == ';' {
        				varName = varName[:last]
    				}

					/* Current variable status:
						queryString has the string to execute "create table" etc.
						varName has the table of databasename
					*/
					valid,message := validQuery(queryString)

					if valid {
						//Hand off query to either the database or table handler
						if message == "db" {
							fmt.Println(databaseHandler(queryString))
						} else {
							fmt.Println(tableHandler(queryString))
						}
					} else {
						fmt.Println("....Invalid query string")
					}
				} else {
					fmt.Println("....Error: Missing ; at end of query")
				}
			}
			fmt.Print("> ")
		}
   }
}

func existsInArr(testStr string, arr []string) bool{
	for i := 0; i < len(arr); i++ {
		if (testStr == arr[i]) {
			return true
		}
	}
	return false
}

func validQuery(queryStr string) (bool,string){
	if existsInArr(queryStr, validDbStringsL) || existsInArr(queryStr, validDbStringsU) {
		return true,"db"
	} else if existsInArr(queryStr, validTableStringsU) || existsInArr(queryStr, validTableStringsL) {
		return true,"table"
	} else {
		return false,"invalid"
	}
}

func databaseHandler(queryStr string) string{
	switch queryStr {
	case "createdatabase":
		return createDatabase()
	case strings.ToUpper("createdatabase"):
		return createDatabase()
	case "dropdatabase":
		return dropDatabase()
	case strings.ToUpper("dropdatabase"):
		return dropDatabase()
	case "use":
		return useDatabase()
	case strings.ToUpper("use"):
		return useDatabase()
	default:
		return "....Error, bad database query string"
	}
}

func tableHandler(queryStr string) string{
	switch queryStr {
	case "createtable":
		return createTable()
	case strings.ToUpper("createtable"):
		return createTable()
	case "droptable":
		return dropTable()
	case strings.ToUpper("droptable"):
		return dropTable()
	case "select*from":
		return selectFrom()
	case strings.ToUpper("select*from"):
		return selectFrom()
	case "altertable":
		return alterTable()
	case strings.ToUpper("altertable"):
		return alterTable()
	default:
		return "....Error, bad table query string"
	}
}

/*     Database functions     */
func createDatabase() string{
	return "Create Database"
}

func dropDatabase() string{
	return "Drop Database"
}

func useDatabase() string{
	return "Use Database"
}

/*     Table functions     */
func createTable() string{
	return "Create Table"
}

func dropTable() string{
	return "Drop Table"
}

func selectFrom() string{
	return "Select From Table"
}

func alterTable() string{
	return "Alter Table"
}
