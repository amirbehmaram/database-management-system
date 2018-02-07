# Database management system by Amir Behmaram

### Installation and setup of GO
--------------------------------

### Compilation instructions
----------------------------
Once go is installed, navigate to the project folder and type
```go build```
This will compile the program

### Running the Program
-----------------------
Once the program is built type to run it:
 ```./database-management-system```

### Supported SQL statements and examples
-----------------------------------------
Currently accepts ALL upper case or lower case versions of the following statements:

#### Databases
1. CREATE DATABASE <db name>;
2. DROP DATABASE <db name>;
3. USE <db name>;

#### Tables
1. CREATE TABLE <tb_name> (<name> <type>, <name2> <type2>);
2. DROP TABLE <tb_name>;
3. SELECT * FROM <tb_name>;
4. ALTER TABLE <tb_name> ADD <name> <type>;
