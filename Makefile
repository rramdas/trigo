MAIN_PACKAGE = cmd/server
BINARIES = bin/
BINARY_NAME = trigo

.PHONY: audit
audit:
	go mod verify
	go vet ./...

.PHONY: clean
clean:
	go clean
	rm -f bin/server

.PHONY: run
run: server
	@./bin/server

.PHONY: server
server:
	@go build -o bin ./cmd/server

.PHONY: test
test:
	@go test -v ./...

.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

.PHONY: deploy
deploy:
	@echo "Deploying...."
	@echo "Building the binary..."
	@GOOS=linux GOARCH=amd64 go build -o $(BINARIES)$(BINARY_NAME) $(MAIN_PACKAGE)
	@echo "Compressing the binary..."
	upx -5 $(BINARIES)$(BINARY_NAME) -o $(BINARIES)$(BINARY_NAME)
	@echo "Uploading the binary..."
	@scp -i ~/.ssh/trigo.pem $(BINARIES)$(BINARY_NAME)
