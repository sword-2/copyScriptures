//Package pipes investigates operating pipes in go.
package pipes

import (
    "fmt"
    "os"
	"os/exec"
)

//This example runs ls command and was adapted from net - https://yifan-online.com/en/km/article/detail/4967
func Ls() {
	// create pipeline
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Println(err)
		return
	}

	// start child process
	cmd := exec.Command("ls")
	cmd.Stdout = w
	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return
	}

	// read data from the pipe
	buf := make([]byte, 1024)
	n, err := r. Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf[:n]))

	// wait for the child process to finish
	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
		return
	}

	// close the pipe
	r.Close()
	w. Close()
}
