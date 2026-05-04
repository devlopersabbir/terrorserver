# Terror Server

A minimal, production-grade domain-based router, reverse proxy, and static file server.

**No plugins. No middleware chains. No complexity. Just routes.**

---

## Features

- **Domain-based routing** — route by hostname or port
- **Reverse proxy** — transparent upstream forwarding
- **Static file server** — with SPA fallback support
- **Runtime watcher** — installer restarts the service when config changes
- **Single binary** — no runtime dependencies
- **Systemd integration** — production-ready service management

---

## Installation

```bash
curl -fsSL https://terror.softvenceomega.com/install.sh | sudo bash
```

The installer will:
1. Pull the latest stable release from GitHub
2. Install to `/usr/local/bin/terror`
3. Create config at `/etc/terror/Runtime`
4. Register and start a `systemd` service

### Environment variables

| Variable        | Default                | Description              |
|-----------------|------------------------|--------------------------|
| `TERROR_CONFIG` | `/etc/terror/Runtime`  | Path to config file      |
| `TERROR_ADDR`   | `:80`                  | Listen address           |
| `TERROR_WEB_ROOT` | `/var/www/terrorserver` | Default welcome site root |
| `TERROR_TEMPLATE_BASE` | `https://terror.softvenceomega.com` | Installer template host |

---

## Config

Edit `/etc/terror/Runtime`. The installer enables a systemd path watcher, so saved changes restart `terror.service` automatically.

```
# Reverse proxy
api.example.com {
    proxy localhost:5000
}

# Static file server (with SPA fallback)
app.example.com {
    root /var/www/html
    file_server
}

# Port-based proxy (any host on :4000)
:4000 {
    proxy localhost:3000
}
```

### Rules

- One block per domain or port
- Each block must have exactly one of `proxy` or `file_server`
- `file_server` requires a `root` directive
- Comments start with `#`
- Unknown directives are a parse error (fail-fast)

---

## Commands

```bash
terror                  # Start the server (default)
terror serve            # Start the server
terror validate         # Validate config without starting
terror status           # Show runtime status
terror version          # Print version
terror help             # Show help
```

---

## Runtime Reload

When installed with `install.sh`, systemd watches `/etc/terror/Runtime`. When a change is saved:

1. `terror.path` detects the file change
2. `terror-restart.service` restarts `terror.service`
3. New domains, ports, proxies, static roots, and SSL listeners are loaded

This restart is intentional because adding or removing listeners such as `:9090` or `:443` requires rebinding ports.

---

## Failure Behavior

| Situation               | Response                          |
|-------------------------|-----------------------------------|
| Invalid config on reload| Keep old config, log error        |
| Proxy upstream down     | `502 Bad Gateway`                 |
| Static file not found   | `404 Not Found`                   |
| Unknown domain/host     | `404 Not Found`                   |

---

## Development

```bash
# Run tests
make test

# Build binary
make build

# Run locally (uses testdata/Runtime, port 8080)
make run

# Cross-compile for Linux/macOS
make build-all
```

---

## Uninstall

```bash
curl -fsSL https://terror.softvenceomega.com/uninstall.sh | sudo bash
```

---

No middleware. No plugin system. No dynamic logic.

---

## License

MIT
