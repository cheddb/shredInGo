package main

import (
    "bufio"
    "math/rand"
    "os"
    "testing"
)

func generateFile(filePath string, fileSize int64){
  file, err := os.Create( filePath )
  defer file.Close()
  check(err)

  w := bufio.NewWriter(file)
  for i:=int64(0); i< fileSize; i++ {
      randomValue := byte(rand.Intn(255))
      check(w.WriteByte(randomValue))
  }

  w.Flush()

}

func TestFileGenerator(t *testing.T) {
    filePath := "testFileGeneration.txt"
    fileRequestedSize := int64(124)
    generateFile(filePath, fileRequestedSize)
    file, err := os.Open( filePath )
    defer file.Close()
    check(err)
    computedFileSize := getFileSize(file)
    if computedFileSize != fileRequestedSize{
      t.Fatalf(`size(generateFile(%s, %v)) = %v, want "%v", error`, filePath, fileRequestedSize, computedFileSize, fileRequestedSize)
    }
}
