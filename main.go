package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	commcid "github.com/filecoin-project/go-fil-commcid"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		commp, err := hex.DecodeString(scanner.Text())
		if err != nil {
			log.Fatalf("Failed to hex decode line (%s): %s", scanner.Text(), err)
			os.Exit(1)
		}

		cid := commcid.PieceCommitmentV1ToCID(commp)
		fmt.Println(cid)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to read line from standard input: %s", err)
		os.Exit(1)
	}
}
