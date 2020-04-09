package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
)

func main() {
  reader := bufio.NewReader(os.Stdin)

  fmt.Println("Who are you?")

  name, _, err := reader.ReadLine()
  if err != nil {
    log.Fatal("Failed to read line from standard input: %w\n", err)
    os.Exit(1)
  }

  fmt.Printf("Hello %s\n", name)
}
