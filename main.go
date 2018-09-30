package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

func main() {
	db, err := NewDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %s", err)
	}
	defer db.Close()

	// input := "atreus62,11,02,1,0,0,26642,0"

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		keyAction, err := NewKeyAction(input)
		if err != nil {
			log.Printf("Error when initializing KeyAction: %s", err)
			continue
		}

		err = db.InsertKeyAction(keyAction, time.Now().Unix())
		if err != nil {
			log.Printf("Could not insert keyaction into the database: %s", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed when reading Stdin: %s", err)
	}
}
