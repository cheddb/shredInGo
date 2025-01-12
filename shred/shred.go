
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
      log.Panic(e)
    }
}

func getFileSize(file *os.File) int64{
  fi, err := file.Stat()
  check(err)
  return fi.Size() // length in bytes for regular files; system-dependent for others
}

func overwriteFileWithRandomValue(file *os.File, fileSize int64){
    ranGen := rand.New(rand.NewSource(time.Now().UnixNano()))

    _, err := file.Seek(0, 0)
    check(err)

    w := bufio.NewWriter(file)
    for i:=int64(0); i< fileSize; i++{
        randomValue := byte(ranGen.Intn(255))
        check(w.WriteByte(randomValue))
    }

    w.Flush()
}

func shred(filePath string) int {
    file, err := os.OpenFile( filePath, os.O_RDWR|os.O_TRUNC, 0755)
    defer file.Close()
    check(err)
    fileSize := getFileSize(file)
    if fileSize > 0{
        for i:=0; i< 3; i++{
            overwriteFileWithRandomValue(file, fileSize)
        }
    }
    os.Remove(filePath)
    return 0
}

func main() {
    helpMessage := fmt.Sprintf(`Usage:
%s [filename]`, os.Args[0])

    if len(os.Args) != 2 {
        log.Println("Wrong number of args.")
        log.Println(helpMessage)
        os.Exit(-1)
    }

    if os.Args[1] == "-h" || os.Args[1] == "--help"{
        log.Println(helpMessage)
        os.Exit(0)
    }

    filepath := os.Args[1]
    os.Exit(shred(filepath))
}
