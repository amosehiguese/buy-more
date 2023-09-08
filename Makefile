MIGRATINS_FOLDER ?= $(PWD)/migrations

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: build
build:
	go build -o bin/

.PHONY: run
run: build
	./bin

.PHONY: mig-create
mig-create:
	migrate create -ext sql -dir migrations -seq create_init_tables

.PHONY: mig-up
mig-up:
	migrate -database "$(DATABASE_URL)" -path $(MIGRATINS_FOLDER) up

.PHONY: mig-down
mig-down:
	migrate -database "$(DATABASE_URL)"  -path $(MIGRATINS_FOLDER) down
