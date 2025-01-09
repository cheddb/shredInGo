package main

import (
    "fmt"
    "os"
)

func main() {
    helpMessage := fmt.Sprintf(`Usage:
%s [filename]`, os.Args[0])
    
    if len(os.Args) != 2 {
        fmt.Println("Too many args.")
        fmt.Println(helpMessage)
        os.Exit(-1)
    }
    
    file := os.Args[1]

    fmt.Println(file)
}
