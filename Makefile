migrate-create:  ### create new migration
	migrate create -ext sql -dir migrations 'users'
.PHONY: migrate-create