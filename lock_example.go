// +build ignore

package main

import (
	"fmt"
	"time"

	"github.com/allan-simon/go-singleinstance"
)

func main() {
	filename := "plop.lock"

	_, err := singleinstance.CreateLockFile(filename)
	if err != nil {
		fmt.Println("An instance already exists")

		pid, err := singleinstance.GetLockFilePid(filename)
		if err != nil {
			fmt.Println("Cannot get PID:", err)
			return
		}

		fmt.Println("Locking PID:", pid)
		return
	}

	fmt.Println("Sleeping...")
	time.Sleep(30 * time.Second)
	fmt.Println("Done")
}
