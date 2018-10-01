package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// DB holds a database connection and methods to interact with it
type DB struct {
	conn *sql.DB
}

// NewDB creates a connection to the SQLite database and returns
// a DB object that can interact with it the database
func NewDB() (*DB, error) {
	conn, err := sql.Open("sqlite3", "/home/taylon/.keylogger/database.db")
	if err != nil {
		return nil, err
	}

	db := &DB{
		conn: conn,
	}

	return db, nil
}

// Close closes the database connection
func (db *DB) Close() {
	db.conn.Close()
}

// InsertKeyAction adds a KeyAction to the database
func (db *DB) InsertKeyAction(keyAction *KeyAction, unixTime int64) error {
	_, err := db.conn.Exec(`INSERT INTO
			key_actions (created_at, keyboard_name, col, row, press, tap_count, tap_interrupted, key_code, layer)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		unixTime,
		keyAction.Keyboard,
		keyAction.Column,
		keyAction.Row,
		keyAction.Press,
		keyAction.TapCount,
		keyAction.TapInterrupted,
		keyAction.KeyCode,
		keyAction.Layer,
	)

	if err != nil {
		return err
	}

	return nil
}
