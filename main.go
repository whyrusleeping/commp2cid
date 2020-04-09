package main

import (
  "bufio"
  "fmt"
  "io"
  "log"
  "os"
)

func ReadLine(reader *bufio.Reader) (string, error) {
  var line, part []byte
  var isPrefix bool = true
  var err error = nil

  for isPrefix && err == nil {
    part, isPrefix, err = reader.ReadLine()
    line = append(line, part...)
  }

  return string(line), err
}

func ConvertCommitmentPieceToCID(commp string) (string) {
  return commp
}

func main() {
  var reader = bufio.NewReader(os.Stdin)
  var commp string
  var err error = nil

  for err == nil {
    commp, err = ReadLine(reader)
    if err == nil {
      cid := ConvertCommitmentPieceToCID(commp)
      fmt.Println(cid)
    }
  }

  if err != io.EOF {
    log.Fatal("Failed to read line from standard input: %w\n", err)
    os.Exit(1)
  }
}
