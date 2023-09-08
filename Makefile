MIGRATIONS_FOLDER = $(PWD)/db/migrations

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: build
build:
	go build -o bin/buy-more

.PHONY: run
run: build
	./bin/buy-more

.PHONY: mig-create
mig-create:
	migrate create -ext sql -dir db/migrations -seq create_init_tables

.PHONY: mig-up
mig-up:
	migrate -database "$(DATABASE_URL)" -path $(MIGRATIONS_FOLDER) up

.PHONY: mig-down
mig-down:
	migrate -database "$(DATABASE_URL)"  -path $(MIGRATIONS_FOLDER) down

.PHONY: mig-force
mig-force:
	migrate -database "$(DATABASE_URL)" -path $(MIGRATIONS_FOLDER) force 1


