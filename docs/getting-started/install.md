# Install

The hosted installer downloads the latest stable Linux release, installs the binary, creates service files, and prepares the default Runtime and welcome site paths.

## Standard Install

```bash
curl -fsSL https://raw.githubusercontent.com/devlopersabbir/terrorserver/main/scripts/install.sh | sudo bash
```

It fetches the stable binary for your host architecture (`amd64` or `arm64`), creates default paths, drops systemd service and watcher units, and starts the daemon.

## Installation Flow

1. Validates root execution (`sudo`).
2. Checks for `curl` or `wget`.
3. Detects OS and architecture (`linux-amd64` or `linux-arm64`).
4. Downloads the release binary from GitHub Releases.
5. Installs binary to `/usr/local/bin/terror`.
6. Downloads `welcome.html` into `/var/www/terrorserver/index.html`.
7. Creates certificate cache at `/var/lib/terror/certs`.
8. Downloads `Runtime` config template into `/etc/terror/Runtime`.
9. Writes systemd service `/etc/systemd/system/terror.service`.
10. Writes systemd path watcher `/etc/systemd/system/terror.path` and restart unit `/etc/systemd/system/terror-restart.service`.
11. Enables and starts the systemd service.

## Installer Overrides

You can control installer behavior using environment variables:

```bash
curl -fsSL https://raw.githubusercontent.com/devlopersabbir/terrorserver/main/scripts/install.sh | TERROR_ADDR=":9090" sudo -E bash
```

Supported overrides:

| Variable | Purpose |
| --- | --- |
| `TERROR_ADDR` | Custom listen port during initial Runtime generation |
| `TERROR_WEB_ROOT` | Custom default web root path |
| `TERROR_CERT_CACHE` | Custom Let's Encrypt certificate directory |
| `TERROR_REPO` | Target GitHub repository (`devlopersabbir/terrorserver`) |
| `TERROR_ASSET` | Specific binary release asset name |
| `TERROR_TEMPLATE_BASE` | `https://raw.githubusercontent.com/devlopersabbir/terrorserver/main/templates` |

## Verify The Install

```bash
terror version
terror status
sudo systemctl status terror
sudo systemctl status terror.path
```

The service should be active, and the path watcher should be enabled so Runtime edits restart the server.

## After Install

1. Point DNS records at the server.
2. Open inbound ports `80` and `443` when using domain routes.
3. Put local apps behind `localhost` ports.
4. Edit `/etc/terror/Runtime`.
5. Run `terror validate` and `terror status`.
