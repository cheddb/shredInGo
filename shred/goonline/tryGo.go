
package main

import (
    "fmt"
    "os"
    "time"
    "math/rand"
    "log"
)

func check(e error) {
    if e != nil {
        log.Fatal(e)
        panic(e)
    }
}

func overwriteFileWithRandomValue(file *File, fileSize int){
    ranGen := rand.New(rand.NewSource(time.Now().UnixNano()))

    err = file.Truncate(0)
    check(err)
    _, err = file.Seek(0, 0)
    check(err)
    
    // randomSuite := make([]byte, fileSize)
    // rand.Read(randomSuite)
    // fmt.Println(token)
    w := bufio.NewWriter(file)
    for i:=0; i< fileSize; i++{
        byte randomValue := ranGen.Intn(255)
        check(w.WriteByte(randomValue))
    }
    
    w.Flush()
}

func main() {
    helpMessage := fmt.Sprintf(`Usage:
%s [filename]`, os.Args[0])
    
    if len(os.Args) != 2 {
        fmt.Println("Too many args.")
        fmt.Println(helpMessage)
        os.Exit(-1)
    }
    
    filepath := os.Args[1]

    fmt.Println(filepath)
    file, err := os.Open( filepath )
    defer file.Close()
    check(err)
    fi, err := file.Stat()
    check(err)
    fmt.Println( fi.Size() )
    fileSize := fi.Size() // length in bytes for regular files; system-dependent for others
    
    for i:=0; i< 3; i++{
        overwriteFileWithRandomValue(file, fileSize)
    }
    
    os.Remove(filepath)
    os.Exit(0)
}
