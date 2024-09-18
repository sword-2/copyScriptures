package pipes

import (
    "bufio"
    "fmt"
    "log"
    "os"
	"sync"
)


//This example adapted from net - https://gosamples.dev/read-user-input/
func Bufio(wg *sync.WaitGroup) {
	//Two arguments
	fmt.Print("Enter some text for program to read by stdin:")

	//if there is a space between 2 words, this method gets both words ...
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
	}
	fmt.Printf("read line: %s-\n", scanner.Text())

	defer wg.Done()
}
