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

	switch commandsArray[0] {
	case "generate":
		fmt.Println(phonescript.GenerateNumbers(commandsArray[1]))
		return
	case "import":
		phonescript.ValidateNumbers(commandsArray[1])
		return
	default:
		fmt.Println("Invalid output")
	}
}
