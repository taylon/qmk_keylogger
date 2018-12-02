package main

import (
	"bufio"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/taylon/qmk_keylogger/db"
	"github.com/taylon/qmk_keylogger/hidlisten"
	"github.com/taylon/qmk_keylogger/keyaction"
)

func watchForSystemSignals(db *db.DB) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		sig := <-sigChan

		switch sig {
		case syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT:
			db.Close()

			log.Printf("exiting after receiving '%s' signal", sig)

			os.Exit(0)
		}
	}()
}

func main() {
	log.SetOutput(os.Stdout)

	listenerDB, err := db.New()
	if err != nil {
		log.Fatalln("could not connect to the database:", err)
	}
	defer listenerDB.Close()

	watchForSystemSignals(listenerDB)

	hidListen, err := hidlisten.New()
	if err != nil {
		log.Fatalln("could not initialize hid_listen:", err)
	}

	hidListenErrChan, err := hidListen.Start()
	if err != nil {
		log.Fatalln("could not start hid_listen:", err)
	}

	// input := "dactyl,11,02,1,0,0,26642,0"
	go func() {
		scanner := bufio.NewScanner(hidListen.StdOutputPipe)
		for scanner.Scan() {
			input := scanner.Text()

			keyAction, err := keyaction.New(input)
			if err != nil {
				log.Printf("error when initializing KeyAction: %s", err)
				continue
			}

			err = listenerDB.InsertKeyAction(keyAction, time.Now().Unix())
			if err != nil {
				log.Printf("could not insert keyaction into the database: %s", err)
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatalln("error when reading hid_listen stdout:", err)
		}
	}()

	err = <-hidListenErrChan
	log.Fatalln("error on hid_listen:", err)
}
