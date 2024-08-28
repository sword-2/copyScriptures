// main.go.

package main

import (
	"fmt"
	"os"
	"tanach-go/tanach"
)


func tanachTest(t *string) {
	tanach.ReadUnmarshalXml(t)
	tanach.PrintTeiHeader()
	tanach.PrintTanach()
	tanach.PrintNotes()
}



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
        case "1":
			tanachTest(&testFile)
        case "2":
			tanach.WriteFile(&testFile)
			var sortCode int // variable declaration
			sortCode = 2 //random code
			tanach.SortStruct(&sortCode)
			tanach.PrintStruct()
			tanach.RemoveFile(&testFile)
        default:
			fmt.Println("Got unexpected arg not in range: ", arg)
    }
}
