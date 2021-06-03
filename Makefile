MAIN_PATH=cmd/calculator/main.go
MODULE=go list -m
GOARCH:=amd64
GOOS:=linux
PRODUCT:=calc
GOENVS:=CGO_ENABLED=0

export GO111MODULE=on

run:
	@echo "Running project"
	go run $(MAIN_PATH)

vendor:
	@echo "Build vendor"
	go mod vendor

tidy:
	go mod tidy

vet:
	go vet ./...

test: vet
	@echo "Running tests"
	go test ./... -v -cover

build: ## build the base bikepark application
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOENVS) go build -v -ldflags \
		"-X main.Version=$(VERSION) $(LDFLAGS)" \
		$(BUILDFLAGS) -o $(PRODUCT) $(MAIN_PATH)

clean:
	rm -rf vendor
	rm $(PRODUCT)

build-docker-image: BUILDFLAGS=""
build-docker-image: test
	docker build \
		--no-cache \
		.

docker-up:
	docker-compose up --force-recreate --build -d

docker-down:
	docker-compose down