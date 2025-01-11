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

func fileExists(filePath string) bool {
  if _, err := os.Stat(filePath); err == nil {
    return true
  }
  return false
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

func TestFileOverwrite(t *testing.T) {
    filePath := "testFileGeneration2.txt"
    fileRequestedSize := int64(124)
    generateFile(filePath, fileRequestedSize)
    file, err := os.Open( filePath )
    defer file.Close()
    check(err)
    overwriteFileWithRandomValue(file, fileRequestedSize)
}

func TestShredEmptyFile(t *testing.T) {
  filePath := "testEmptyFile.txt"
  fileRequestedSize := int64(0)
  generateFile(filePath, fileRequestedSize)

  if shred(filePath) != 0{
    t.Fatalf(`shred failed with an empty file`)
  }
  if fileExists(filePath){
    t.Fatalf(`shred failed to remove an empty file`)
  }

}
func TestShredNotAFile(t *testing.T) {
  defer func() {
   if r := recover(); r != nil {
    t.Log("Test passed, Nothing has been done!")
   }
  }()
  filePath := "notAFile.txt"
  if fileExists(filePath){
    t.Fatalf(`%s exists while it shouldn't to run this test`, filePath)
  }
  err := shred(filePath)
  if err == 0{
    t.Errorf("did not panic")
  }
}
func TestShredReadOnlyFile(t *testing.T) {
  defer func() {
   if r := recover(); r != nil {
     t.Log("Test passed, Nothing has been done!")
   }
  }()
  filePath := "readOnlyFile.txt"
  fileRequestedSize := int64(0)
  if fileExists(filePath){
    err := os.Chmod(filePath, 0774)
    if err != nil {
      t.Fatalf(`Failed to make the file writable: %s`, err)
    }
  }
  generateFile(filePath, fileRequestedSize)
  err := os.Chmod(filePath, 0444)
  if err != nil {
    t.Fatalf(`Failed to make the file read only: %s`, err)
  }

  if shred(filePath) != 0{
    t.Fatalf(`shred failed with an empty file`)
  }
  if fileExists(filePath){
    t.Fatalf(`shred failed to remove an empty file`)
  }

}
// func TestShredLargeFile(t *testing.T) {
// }
// func TestShredDirectory(t *testing.T) {
// }
// func TestShredSymbolicFile(t *testing.T) {
// }
// func TestShredConcurrentAccess(t *testing.T) {
// }
// func TestShredHiddenFile(t *testing.T) {
// }
// func TestReaccessData(t *testing.T) {
// }
