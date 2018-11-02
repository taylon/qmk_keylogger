.PHONY: create_db test db all build

PREFIX ?= /usr/bin

KEYLOGGER_FILES_PATH=$(HOME)/.keylogger
DB_FILE=$(KEYLOGGER_FILES_PATH)/database.db
INIT_SQL_FILE=db.sql

all: build

create_db:
	@cat $(INIT_SQL_FILE) | sqlite3 $(DB_FILE)

db:
	@sqlite3 $(DB_FILE)

test:
	@go test -v -cover -json | tparse -all

build:
	@go build -o qmk_keylogger

install: build
	cp qmk_keylogger $(PREFIX)/