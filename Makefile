.PHONY: create_db

KEYLOGGER_FILES_PATH=$(HOME)/.keylogger
DB_FILE=$(KEYLOGGER_FILES_PATH)/database.db
INIT_SQL_FILE=db.sql

create_db:
	@cat $(INIT_SQL_FILE) | sqlite3 $(DB_FILE)

db:
	@sqlite3 $(DB_FILE)

test:
	@go test -v -cover

run:
	@go build
	@sudo ./hid_listen | ./qmk_keylogger >> $(KEYLOGGER_FILES_PATH)/logs
