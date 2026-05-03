.PHONY: build test clean install lint run dev

BINARY   := terror
CMD      := ./cmd/terror
LDFLAGS  := -ldflags="-s -w"
GOFLAGS  :=

# ─── Build ────────────────────────────────────────────────────────────────────

build:
	go build $(LDFLAGS) -o $(BINARY) $(CMD)

build-linux:
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY)-linux-amd64 $(CMD)

build-all: build-linux
	GOOS=linux GOARCH=arm64  go build $(LDFLAGS) -o $(BINARY)-linux-arm64  $(CMD)
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY)-darwin-amd64 $(CMD)
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(BINARY)-darwin-arm64 $(CMD)

# ─── Test ─────────────────────────────────────────────────────────────────────

test:
	go test ./... -v -race

test-short:
	go test ./... -short

bench:
	go test ./... -bench=. -benchmem

# ─── Dev ──────────────────────────────────────────────────────────────────────

run: build
	TERROR_ADDR=":8080" TERROR_CONFIG="./testdata/Runtime" ./$(BINARY)

# ─── Code quality ─────────────────────────────────────────────────────────────

lint:
	@command -v golangci-lint >/dev/null 2>&1 && golangci-lint run ./... || \
	  echo "golangci-lint not installed: https://golangci-lint.run/usage/install/"

vet:
	go vet ./...

# ─── Cleanup ──────────────────────────────────────────────────────────────────

clean:
	rm -f $(BINARY) $(BINARY)-*

# ─── Install (requires root) ──────────────────────────────────────────────────

install: build
	sudo bash scripts/install.sh

uninstall:
	sudo bash scripts/uninstall.sh
