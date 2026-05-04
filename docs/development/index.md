# Development

Use this page when working on Terror Server itself.

## Requirements

- Go 1.25 or newer.
- Linux for installer and systemd testing.

## Tests

```bash
make test
```

Useful direct checks:

```bash
go test ./...
go vet ./...
bash -n scripts/install.sh
bash -n scripts/uninstall.sh
```

## Build

Build locally:

```bash
make build
```

Build Linux release assets:

```bash
make build-linux
make build-all
```

## Run Manually

Run with a custom Runtime:

```bash
TERROR_CONFIG=/path/to/Runtime TERROR_ADDR=:8080 go run ./cmd/terror serve
```

This is useful for local routing experiments without touching the installed production Runtime.
