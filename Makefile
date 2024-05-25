include .env
export

# LOCAL_BIN:=$(CURDIR)/bin
# PATH:=$(LOCAL_BIN):$(PATH)

migrate-create:  ### create new migration
	migrate create -ext sql -dir migrations 'users'
.PHONY: migrate-create

run: 
	go run cmd/app/main.go
.PHONY: run