package pipes

import (
    "fmt"
	"sync"
    //"os"
	//"os/exec"
)


//This example only works if the quanity of input arguments is known and fixed.
func FixedArgs(wg *sync.WaitGroup) {
	//Two arguments
	fmt.Print("Enter some text for program to read by stdin:")

	//if there is a space between 2 words, this method gets both words ...
		var s1 string
		var s2 string
		fmt.Scanln(&s1,&s2)
		fmt.Println(s1)
		fmt.Println(s2)

	defer wg.Done()
}
