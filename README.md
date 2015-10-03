# go-singleinstance
cross plateform library to have only one instance of a software (based on python's trendo)


## Usage

```
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

```

if you try to launch it twice, the second instance will fail


## Thanks

For the python library trendo, from which I've shamelessly adapted the code

## Contribution

Don't be afraid if it says "last commit 2 years ago", this library is made to be small
and simple so it's unlikely it changes after some times, however I'm pretty reactive
on github overall, so feel free to use issues to ask question, propose patch etc. :)
