package main

import (
    "fmt"
    "os"
)

func main() {
    for idx, arg := range os.Args {
        fmt.Printf("Args[%d]: %s\n", idx, arg)
    }
}
