# Terror Server

Terror Server is a small production-oriented HTTP router, reverse proxy, static file server, and automatic TLS gateway packaged as a single Go binary.

It is designed for simple Linux servers where you want Caddy-like ergonomics without a large configuration surface: define routes in one Runtime file, point domains or ports at upstream apps, serve static sites, and let the service manage HTTP, HTTPS, and systemd operation.

## Features

- Domain-based routing by hostname.
- Port-based routing such as `:9090`.
- Reverse proxy support for local upstream apps.
- Jenkins-friendly `X-Forwarded-*` proxy headers.
- Static file serving with index fallback for single-page apps.
- Automatic Let's Encrypt certificates for domain routes.
- Optional HTTP to HTTPS redirects for domain routes.
- Default welcome page on port `80`.
- Runtime health checks with `terror status`.
- Systemd service hardening and automatic restart on Runtime changes.
- Stable release workflow with Linux `amd64` and `arm64` binaries.
- Hosted installer and uninstaller assets for simple curl-based installation.

## Quick Start

Install the latest stable release:

```bash
curl -fsSL https://terror.softvenceomega.com/install.sh | sudo bash
```

Check the service:

```bash
terror status
```

Edit the Runtime config:

```bash
sudo vim /etc/terror/Runtime
```

The installer enables a systemd path watcher, so saving `/etc/terror/Runtime` restarts `terror.service` automatically. This allows new domains, removed routes, new listener ports, and TLS listeners to be picked up without manual `systemctl restart`.

## Default Paths

| Path                                         | Purpose                                   |
| -------------------------------------------- | ----------------------------------------- |
| `/usr/local/bin/terror`                      | Installed binary                          |
| `/etc/terror/Runtime`                        | Main Runtime configuration                |
| `/var/www/terrorserver`                      | Default static welcome site root          |
| `/var/www/terrorserver/index.html`           | Default welcome page                      |
| `/var/lib/terror/certs`                      | Let's Encrypt certificate cache           |
| `/etc/systemd/system/terror.service`         | Main systemd service                      |
| `/etc/systemd/system/terror.path`            | Runtime change watcher                    |
| `/etc/systemd/system/terror-restart.service` | Restart helper triggered by `terror.path` |

## Runtime Configuration

Terror Server uses a compact Caddy-style block format. Each block starts with either a domain name or a port, and must contain exactly one route handler.

```txt
:80 {
    root /var/www/terrorserver
    file_server
}

app.example.com {
    proxy localhost:4000
}

:9090 {
    proxy localhost:4000
}

static.example.com {
    root /var/www/html
    file_server
}
```

### Reverse Proxy

Proxy a domain to an app running locally:

```txt
api.example.com {
    proxy localhost:3000
}
```

Proxy any request received on a port:

```txt
:9090 {
    proxy localhost:4000
}
```

### Static Sites

Serve files from a directory:

```txt
example.com {
    root /var/www/example
    file_server
}
```

If a requested static path does not exist, Terror Server falls back to `/`. This makes frontend apps with client-side routing work cleanly.

### Config Rules

- One block per domain or port.
- A block can use `proxy` or `file_server`, not both.
- `file_server` requires a `root` directive.
- `proxy` requires an upstream address.
- Comments start with `#`.
- Unknown directives fail validation.
- Domains are matched case-insensitively.

Validate before saving a production change:

```bash
terror validate
```

## HTTPS and Domains

For domain routes, Terror Server enables automatic Let's Encrypt TLS by default.

Requirements:

- The domain must resolve to the server's public IP.
- Inbound ports `80` and `443` must be open.
- No other service can already bind ports `80` or `443`.
- The installed service must have write access to `/var/lib/terror/certs`.

When TLS is enabled:

- Port `443` serves HTTPS for domain routes.
- Port `80` handles ACME HTTP challenges.
- Normal HTTP domain traffic stays available by default while HTTPS is served on port `443`.
- HTTP to HTTPS redirects can be enabled after certificate issuance is confirmed.
- IP-based and port-based routes such as `:9090` stay on HTTP.

Disable automatic TLS only when needed:

```bash
sudo systemctl edit terror
```

Add:

```ini
[Service]
Environment=TERROR_AUTO_TLS=false
```

Then reload:

```bash
sudo systemctl daemon-reload
sudo systemctl restart terror
```

Enable forced HTTPS redirects only after `https://your-domain` is confirmed healthy:

```ini
[Service]
Environment=TERROR_HTTPS_REDIRECT=true
```

## Commands

| Command           | Description                                                   |
| ----------------- | ------------------------------------------------------------- |
| `terror`          | Start the server                                              |
| `terror serve`    | Start the server                                              |
| `terror start`    | Start the server                                              |
| `terror validate` | Validate Runtime config                                       |
| `terror status`   | Show service, listener, DNS, TLS, static, and upstream health |
| `terror update`   | Pull and install the latest stable release                    |
| `terror upgrade`  | Alias for `terror update`                                     |
| `terror version`  | Print binary version                                          |
| `terror help`     | Show command help                                             |

Short aliases:

| Alias       | Command           |
| ----------- | ----------------- |
| `terror s`  | `terror serve`    |
| `terror st` | `terror status`   |
| `terror v`  | `terror validate` |
| `terror u`  | `terror update`   |

## Status Output

`terror status` checks the local installation and route health:

```txt
terrorserver status
-------------------------------------
ok config: /etc/terror/Runtime
ok listen: :80
ok routes: 3 configured
ok service: terror is active
ok watcher: terror.path is active
ok listener: :80 is accepting local connections
ok listener: :443 is accepting local connections
ok ssl: automatic Let's Encrypt SSL enabled
ok ssl: domain HTTP stays available; set TERROR_HTTPS_REDIRECT=true to force HTTPS

routes
ok :80 -> static /var/www/terrorserver (serving static files)
ok app.example.com -> proxy localhost:4000 (upstream reachable)
  ok dns: app.example.com -> 203.0.113.10
ok :9090 -> proxy localhost:4000 (upstream reachable)
```

Use it after every Runtime change. It is intentionally practical: it checks the config file, systemd service, Runtime watcher, expected listeners, DNS resolution, static roots, proxy upstream reachability, and HTTPS availability.

## Updates

Update an installed server to the newest stable release:

```bash
terror update
```

This runs the hosted installer again:

```bash
https://terror.softvenceomega.com/install.sh
```

The installer downloads the latest GitHub release asset, replaces `/usr/local/bin/terror`, refreshes the systemd service files, and keeps existing Runtime config and welcome files when they already exist.

## Uninstall

Interactive uninstall:

```bash
curl -fsSL https://terror.softvenceomega.com/uninstall.sh | sudo bash
```

Non-interactive uninstall:

```bash
curl -fsSL https://terror.softvenceomega.com/uninstall.sh | sudo bash -s -- --yes
```

The uninstaller stops and disables the systemd service, removes the binary and watcher units, then asks before deleting config, welcome site, and certificate cache directories.

## Environment Variables

| Variable               | Default                                        | Used by                | Description                                           |
| ---------------------- | ---------------------------------------------- | ---------------------- | ----------------------------------------------------- |
| `TERROR_CONFIG`        | `/etc/terror/Runtime`                          | Binary, service        | Runtime config path                                   |
| `TERROR_ADDR`          | `:80`                                          | Binary, installer      | Default listen address                                |
| `TERROR_CERT_CACHE`    | `/var/lib/terror/certs`                        | Binary, installer      | Let's Encrypt certificate cache                       |
| `TERROR_AUTO_TLS`      | enabled                                        | Binary                 | Set to `false`, `0`, or `no` to disable automatic TLS |
| `TERROR_HTTPS_REDIRECT` | disabled                                      | Binary                 | Set to `true`, `1`, or `yes` to redirect domain HTTP traffic to HTTPS |
| `TERROR_INSTALL_URL`   | `https://terror.softvenceomega.com/install.sh` | `terror update`        | Installer URL used by update command                  |
| `TERROR_WEB_ROOT`      | `/var/www/terrorserver`                        | Installer, uninstaller | Default welcome page root                             |
| `TERROR_TEMPLATE_BASE` | `https://terror.softvenceomega.com`            | Installer              | Host for `Runtime` and `welcome.html` templates       |
| `TERROR_REPO`          | `devlopersabbir/terrorserver`                  | Installer              | GitHub repository used for release downloads          |
| `TERROR_ASSET`         | auto-detected                                  | Installer              | Release asset override                                |
| `TERROR_CHANNEL`       | `stable`                                       | Installer              | Display channel label                                 |

## Systemd Operation

Common production commands:

```bash
sudo systemctl status terror
sudo journalctl -u terror -n 100 --no-pager
sudo journalctl -uf terror
sudo systemctl restart terror
sudo systemctl status terror.path
```

The installed service uses:

- `Restart=always`
- `CAP_NET_BIND_SERVICE` for binding low ports.
- `ProtectSystem=strict`
- `ReadWritePaths=/etc/terror /var/lib/terror`
- `PrivateTmp=yes`
- `NoNewPrivileges=yes`

## Troubleshooting

### Domain does not route

Check DNS and local host routing:

```bash
nslookup app.example.com
curl -H "Host: app.example.com" http://127.0.0.1
terror status
```

If the local curl works but the browser does not, the issue is usually DNS propagation, firewall rules, cloud security group rules, or another service in front of the server.

### HTTPS shows as not secure

Check the TLS listener and service logs:

```bash
terror status
sudo journalctl -u terror -n 120 --no-pager
curl -vk https://app.example.com
```

Let's Encrypt needs public access to port `80` for the HTTP challenge and port `443` for HTTPS traffic.

### New port route does not work

Confirm the Runtime file was saved and the path watcher is active:

```bash
sudo systemctl status terror.path
sudo journalctl -u terror -n 80 --no-pager
terror status
```

Adding or removing listener ports requires a process restart because ports must be rebound. The installer configures this restart automatically through `terror.path`.

If `terror status` shows `warn watcher`, reinstall or update so the latest systemd watcher files are written:

```bash
terror update
sudo systemctl status terror.path
```

### Jenkins says reverse proxy setup is broken

Jenkins expects reverse proxies to pass the original request information. Terror Server sends:

- `Host`
- `X-Forwarded-Host`
- `X-Forwarded-Proto`
- `X-Forwarded-Port`
- `X-Forwarded-For`
- `X-Real-IP`

Use a normal domain route:

```txt
jenkins.example.com {
    proxy localhost:8080
}
```

Then check:

```bash
curl -H "Host: jenkins.example.com" http://127.0.0.1/login
terror status
```

### Upstream is unreachable

Check the app directly from the server:

```bash
curl http://127.0.0.1:4000
terror status
```

Use `localhost:4000` or `127.0.0.1:4000` in Runtime when the app is running on the same machine.

## Development

Requirements:

- Go 1.25 or newer.
- Linux for installer/systemd testing.

Run tests:

```bash
make test
```

Build locally:

```bash
make build
```

Build Linux release assets:

```bash
make build-linux
make build-all
```

Run manually with a custom Runtime:

```bash
TERROR_CONFIG=/path/to/Runtime TERROR_ADDR=:8080 go run ./cmd/terror serve
```

Useful direct checks:

```bash
go test ./...
go vet ./...
bash -n scripts/install.sh
bash -n scripts/uninstall.sh
```

## Release Flow

Merging to `main` triggers the stable release workflow:

1. Finds the latest `v*-stable` tag.
2. Increments the patch version.
3. Builds Linux `amd64` and `arm64` binaries.
4. Injects the release tag into `terror version`.
5. Generates release notes from commit messages.
6. Publishes the GitHub release as latest.

Installer assets are deployed separately. After a successful stable release, `.github/workflows/deploy-installers.yml` copies changed installer assets to:

```txt
/var/www/terrorserver
```

## License

MIT
