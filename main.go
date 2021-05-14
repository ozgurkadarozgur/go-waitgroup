package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	// wg is here to handle goroutines.
	wg sync.WaitGroup
)

// main function of the app.
func main() {

	// texts are ready to be printed.
	textArr := []string{"a1", "a2", "a3", "a4"}

	// inform wg to wait operations as much as length of textArr.
	wg.Add(len(textArr))

	for _, text := range textArr {
		// call print function as goroutine.
		go print(text)
	}

	// hold on main function until receive done signal as much as length of textArr.
	wg.Wait()
}

// the function prints the given string arg.
func print(text string) {
	// defer keyword does mean; call the code when end of the function even if the function breaks.
	// done function informs waitGroup about the operation is over.
	defer wg.Done()

	// print operations start.
	fmt.Println("waiting to be printed:", text)
	time.Sleep(time.Second * 2)
	fmt.Println("printed:", text)
	// print operations end.
}
