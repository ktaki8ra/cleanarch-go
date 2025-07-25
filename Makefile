TARGET_APP = cleanarch-go
TARGET_MIGRATE = db-migrate

build:
	go build -o $(TARGET_APP) main.go

dbmigrate:
	go build -o $(TARGET_MIGRATE) ./external/migrate/migrate.go

test:
	go test -race -v ./...

clean:
	rm -rf $(TARGET_APP) $(TARGET_MIGRATE)
