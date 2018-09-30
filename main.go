package main

import (
	"log"
	"time"
)

func main() {
	db, err := NewDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %s", err)
	}
	defer db.Close()

	// reader := bufio.NewReader(os.Stdin)

	// for {
	// text, _ := reader.ReadString('\n')
	// atreus62: col=11, row=02, pressed=1, tap_count=0, tap_interrupted=0, keycode=26642, layer=0
	// fmt.Printf("From keylogger: %s", input)
	// }
	input := "atreus62,11,02,1,0,0,26642,0"

	keyAction, err := NewKeyAction(input)
	if err != nil {
		log.Printf("Error when initializing KeyAction: %s", err)
		return
	}

	err = db.InsertKeyAction(keyAction, time.Now().Unix())
	if err != nil {
		log.Printf("Could not insert keyaction into the database: %s", err)
	}
}
