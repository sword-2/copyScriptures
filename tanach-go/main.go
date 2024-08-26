// main.go.

package main

import (
	"fmt"
	//"tanach"
	"tanach-go/tanach"
)


func main() {
	//check if there are any command line arguments
	 fmt.Println("main.go - number arges=", len(os.Args), "; the args=", os.Args)
	//argsWithProg := os.Args
    //argsWithoutProg := os.Args[1:]

	tst := "Books/Genesis.xml"
	//tst := "test.xml"

	fmt.Println("go's main function about to try plantMain()\n")
	fmt.Println("go's main function after plantMain()\n")

	tanach.ReadUnmarshalXml(&tst)
	tanach.PrintTeiHeader()
	tanach.PrintTanach()
	tanach.PrintNotes()
	tanach.WriteFile(&tst)
	var sortCode int // variable declaration
	sortCode = 2 //1 sort on agency name, 2 sort on legal reference.
	tanach.SortStruct(&sortCode)
	tanach.PrintStruct()
	tanach.RemoveFile(&tst)
	//tanach.RemoveOldL14(&pathL14) //if an old .l14 is present, remove it
	//pathTanach := "Tanach/Amons.xml"

	//sortCode = 2 //1 sort on agency name, 2 sort on legal reference.
	//tanach.SortStruct(&sortCode)
	//tanach.PrintStruct()
	//tanach.WriteL14(&pathL14)
}
