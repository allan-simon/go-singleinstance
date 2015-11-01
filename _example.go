package main

import (
	"fmt"
	"time"
)

func main() {

	_, err := CreateLockFile("plop.lock")
	if err != nil {
		fmt.Println("an instance already exists")
		return
	}

	time.Sleep(10 * time.Second)
	fmt.Println("end")
}
