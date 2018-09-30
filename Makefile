.PHONY: create_db

DB_FILE=keylogger.db
INIT_SQL_FILE=db.sql

create_db:
	@cat $(INIT_SQL_FILE) | sqlite3 $(DB_FILE)

test:
	@go test -v -cover