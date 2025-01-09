
package main

import (
    "bufio"
    "fmt"
    "log"
    "math/rand"
    "os"
    "time"
)

func check(e error) {
    if e != nil {
        log.Fatal(e)
        panic(e)
    }
}

func getFileSize(file *os.File) int64{
  fi, err := file.Stat()
  check(err)
  fmt.Println( fi.Size() )
  return fi.Size() // length in bytes for regular files; system-dependent for others
}

func overwriteFileWithRandomValue(file *os.File, fileSize int64){
    ranGen := rand.New(rand.NewSource(time.Now().UnixNano()))

    err := file.Truncate(0)
    check(err)
    _, err = file.Seek(0, 0)
    check(err)

    // randomSuite := make([]byte, fileSize)
    // rand.Read(randomSuite)
    // fmt.Println(token)
    w := bufio.NewWriter(file)
    for i:=int64(0); i< fileSize; i++{
        randomValue := byte(ranGen.Intn(255))
        check(w.WriteByte(randomValue))
    }

    w.Flush()
}

func shred(filePath string) int {

    fmt.Println(filePath)
    file, err := os.OpenFile( filePath, os.O_RDWR|os.O_TRUNC, 0755)
    defer file.Close()
    check(err)
    fileSize := getFileSize(file)

    for i:=0; i< 3; i++{
        overwriteFileWithRandomValue(file, fileSize)
    }

    os.Remove(filePath)
    return 0
}

func main() {
    helpMessage := fmt.Sprintf(`Usage:
%s [filename]`, os.Args[0])

    if len(os.Args) != 2 {
        fmt.Println("Wrong number of args.")
        fmt.Println(helpMessage)
        os.Exit(-1)
    }

    filepath := os.Args[1]
    os.Exit(shred(filepath))
}
