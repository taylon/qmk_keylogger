package main

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type dbTestSuite struct {
	suite.Suite

	db      *DB
	sqlmock sqlmock.Sqlmock
}

func TestRunDBSuite(t *testing.T) {
	suite.Run(t, new(dbTestSuite))
}

func (s *dbTestSuite) SetupTest() {
	conn, mock, err := sqlmock.New()
	if err != nil {
		s.T().Fatalf("Error initializing test database: %s", err)
	}

	s.sqlmock = mock
	s.db = &DB{
		conn: conn,
	}
}

func (s *dbTestSuite) TearDownTest() {
	s.db.Close()
}

func (s *dbTestSuite) TestInsertKeyAction() {
	keyAction := &KeyAction{}
	unixTime := time.Now().Unix()

	s.sqlmock.ExpectExec("INSERT INTO key_actions").
		WithArgs(
			unixTime,
			keyAction.Keyboard,
			keyAction.Column,
			keyAction.Row,
			keyAction.Press,
			keyAction.TapCount,
			keyAction.TapInterrupted,
			keyAction.KeyCode,
			keyAction.Layer,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := s.db.InsertKeyAction(keyAction, unixTime)

	s.NoError(err)
	s.NoError(s.sqlmock.ExpectationsWereMet())
}

func (s *dbTestSuite) TestInsertKeyActionReturnsErrorWhenInsertFails() {
	s.sqlmock.ExpectExec("INSERT INTO key_actions").WillReturnError(errors.New("error"))

	err := s.db.InsertKeyAction(&KeyAction{}, time.Now().Unix())

	s.Error(err)
	s.NoError(s.sqlmock.ExpectationsWereMet())
}
