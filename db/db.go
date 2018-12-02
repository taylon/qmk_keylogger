package db

import (
	"database/sql"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"

	"github.com/taylon/qmk_keylogger/keyaction"
)

// DB holds a database connection and methods to interact with it
type DB struct {
	conn *sql.DB
}

// New creates a connection to the SQLite database and returns
// a DB object that can interact with it the database
func New() (*DB, error) {
	dbPath := path.Join(os.Getenv("HOME"), "/.keylogger/database.db")
	// dbPath := "./dev_tests/database.db"

	conn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	db := &DB{
		conn: conn,
	}

	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		err = db.createTablesAndIndexes()
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

// InsertKeyAction adds a KeyAction to the database
func (db *DB) InsertKeyAction(keyAction *keyaction.KeyAction, unixTime int64) error {
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

func (db *DB) WithTransaction(fn func(*sql.Tx) error) error {
	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = fn(tx)

	return err
}

// createTablesAndIndexes will create the key_actions table and it's indexes
func (db *DB) createTablesAndIndexes() error {
	return db.WithTransaction(func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			CREATE TABLE key_actions (
  				created_at INTEGER,
  				keyboard_name TEXT,
  				col INTEGER,
  				row INTEGER,
  				press INTEGER,
  				tap_count INTEGER,
  				tap_interrupted INTEGER,
  				key_code INTEGER,
  				layer INTEGER
			)`); err != nil {
			return err
		}

		if _, err := tx.Exec(
			"CREATE INDEX key_actions_time_index ON key_actions(created_at)"); err != nil {
			return err
		}

		if _, err := tx.Exec(
			"CREATE INDEX key_actions_keyboard_index ON key_actions(keyboard_name)"); err != nil {
			return err
		}

		return nil
	})
}

// Close closes the database connection
func (db *DB) Close() {
	_ = db.conn.Close()
}
