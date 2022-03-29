build:
	go build -o  bin/main  cmd/main/main.go

run:
	go run cmd/main/main.go

clean:
	rm -rf $(CURDIR)/bin

test:
	go test ./...