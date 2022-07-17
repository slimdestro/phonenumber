# (+)phonescript
#### Bulk Generate && Validate US Phone numbers

[![N|Solid](https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/60px-Go_Logo_Blue.svg.png)](https://dev.to/slimdestro)
 
## Installation

Install the package by doing get get:

```sh
go get github.com/slimdestro/phonenumber/phonescript
```

Two Methods that is exported and can be used in your module are:

```sh
GenerateNumbers(10000) 
ValidateNumbers("/path/to/csv_10K_records")  
```

## Example

```sh
package main 

import (
	"flag"
	"fmt" 
	"github.com/slimdestro/phonenumber/phonescript"
)
var generateNumbersFlag string

func main(){
    flag.StringVar(&generateNumbersFlag, "o", "import", "100")
    flag.Parse()
    commandsArray := flag.Args()
	
	// show Hints instead Index out of range 
	if len(commandsArray) == 0{
		fmt.Println("Forgort ? (generate | import)")
		return
	}

	switch commandsArray[0] {
	case "generate":
		phonescript.GenerateNumbers(commandsArray[1])
		return
	case "import":
		phonescript.ValidateNumbers(commandsArray[1])
		return
	default:
		fmt.Println("Invalid output")
	}
}
 
```


## Author

[slimdestro(Mukul Mishra)](https://linktr.ee/slimdestro)
