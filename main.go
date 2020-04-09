package main

import (
  "bufio"
  "fmt"
  "io"
  "log"
  "os"
  commcid "github.com/filecoin-project/go-fil-commcid"
)

func ReadLine(reader *bufio.Reader) ([]byte, error) {
  var line, part []byte
  var isPrefix bool = true
  var err error = nil

  for isPrefix && err == nil {
    part, isPrefix, err = reader.ReadLine()
    line = append(line, part...)
  }

  return line, err
}

func main() {
  var reader = bufio.NewReader(os.Stdin)
  var commp []byte
  var err error = nil

  for err == nil {
    commp, err = ReadLine(reader)
    if err == nil {
      cid := commcid.PieceCommitmentV1ToCID(commp)
      fmt.Println(cid)
    }
  }

  if err != io.EOF {
    log.Fatal("Failed to read line from standard input: %w\n", err)
    os.Exit(1)
  }
}
