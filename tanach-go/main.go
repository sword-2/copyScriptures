// main.go.

package main

import (
	"fmt"
	"os"
	"tanach-go/tanach"
	"tanach-go/pipes"
	"sync"
)


func tanachTest(t *string) {
	tanach.ReadUnmarshalXml(t)
	tanach.PrintTeiHeader()
	tanach.PrintTanach()
	tanach.PrintNotes()
}

//The program is capable of completing early and this variable tries to introduce delay and prevent that.
var wg sync.WaitGroup

func main() {
	fmt.Println("go's main function ran\n")
	testFile := "Books/Genesis.xml" //initial file to load

	//check if there are any command line arguments
	var arg string = ""
	if len(os.Args) != 2 {
		fmt.Println("main.go - expected 2 arguments, received:", len(os.Args), "; the args were =", os.Args)
		return
	} else {
		//argsWithProg := os.Args
		arg = os.Args[1]
	 }

	switch arg {
        case "1":			//case 1 - test to read a file and print selected tags
			tanachTest(&testFile)
        case "2":			//case 2 to be determined
			tanach.WriteFile(&testFile)
			var sortCode int // variable declaration
			sortCode = 2 //random code
			tanach.SortStruct(&sortCode)
			tanach.PrintStruct()
			tanach.RemoveFile(&testFile)
		case "3": //case 3 learning to use a pipe in go
			//Add a new entry to the waitgroup
			wg.Add(1)
			//pipes.Ls() //Calls a shell command and gets result from pipe
			//go pipes.FixedArgs(&wg) //reads 2 args from stdin
			go pipes.Bufio(&wg) //reads whole line upto newline character
        default:
			fmt.Println("Got unexpected arg not in range: ", arg)
    }
	//Wait until everything is done
    wg.Wait()
}
