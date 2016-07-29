// +build ignore

package main

import (
	"fmt"
	"time"

	"github.com/allan-simon/go-singleinstance"
)

func main() {
	_, err := singleinstance.CreateLockFile("plop.lock")
	if err != nil {
		fmt.Println("An instance already exists")
		return
	}

	fmt.Println("Sleeping...")
	time.Sleep(30 * time.Second)
	fmt.Println("Done")
}
