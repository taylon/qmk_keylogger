.PHONY: create_db test db all build-keylogger run-keylogger

PREFIX ?= /usr/bin

LISTENER_BIN=qmk_keylogger
KEYLOGGER_FILES_PATH=$(HOME)/.keylogger
DB_FILE=$(KEYLOGGER_FILES_PATH)/database.db
INIT_SQL_FILE=./db/db.sql

create_db:
	@cat $(INIT_SQL_FILE) | sqlite3 $(DB_FILE)
all: build-keylogger

db:
	@sqlite3 $(DB_FILE)

test:
	@go test -v -cover -json ./... | tparse -all

build-keylogger:
	@cd keylogger; go build -o $(LISTENER_BIN)

run-keylogger: build-keylogger
	@keylogger/$(LISTENER_BIN)

install: build-keylogger
	cp keylogger/$(LISTENER_BIN) $(PREFIX)/
