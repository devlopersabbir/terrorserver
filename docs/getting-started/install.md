# Installation

Terror Server can be installed directly onto any Linux server via the hosted one-line installer, or run in containerized environments using **Docker** and **Docker Compose**.

## Method 1: Linux Binary Install (Recommended)

The hosted installer automatically downloads the latest compiled binary for your architecture (`amd64` or `arm64`), creates required directories, configures `systemd` service and path watcher units, and starts Terror Server.

```bash
curl -fsSL https://raw.githubusercontent.com/devlopersabbir/terrorserver/main/scripts/install.sh | sudo bash
```

### Installation Flow

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

### Installer Overrides

You can customize installer behavior with environment variables:

```bash
curl -fsSL https://raw.githubusercontent.com/devlopersabbir/terrorserver/main/scripts/install.sh | TERROR_ADDR=":9090" sudo -E bash
```

| Variable | Default | Purpose |
| --- | --- | --- |
| `TERROR_ADDR` | `:80` | Custom listen port during initial Runtime generation |
| `TERROR_WEB_ROOT` | `/var/www/terrorserver` | Custom default web root path |
| `TERROR_CERT_CACHE` | `/var/lib/terror/certs` | Custom Let's Encrypt certificate directory |
| `TERROR_REPO` | `devlopersabbir/terrorserver` | Target GitHub repository |
| `TERROR_ASSET` | Auto-detected | Specific binary release asset name |

---

## Method 2: Docker Container

Terror Server publishes official container images to GitHub Container Registry (`ghcr.io`).

### Quick Start (`docker run`)

Run Terror Server as a background daemon with ports `80` and `443` bound:

```bash
docker run -d \
  --name terror \
  --restart unless-stopped \
  -p 80:80 \
  -p 443:443 \
  -v $(pwd)/Runtime:/etc/terror/Runtime:ro \
  -v terror_certs:/var/lib/terror/certs \
  -v $(pwd)/html:/var/www/html:ro \
  ghcr.io/devlopersabbir/terrorserver:latest
```

::: tip Persistent TLS Certificates
Always mount a named volume or persistent host directory to `/var/lib/terror/certs`. This ensures Let's Encrypt TLS certificates persist across container upgrades and avoids hitting ACME rate limits.
:::

---

### Docker Compose

For multi-container setups and production deployments, create a `docker-compose.yml` file:

```yaml
services:
  terror:
    image: ghcr.io/devlopersabbir/terrorserver:latest
    container_name: terror
    restart: unless-stopped
    ports:
      - "80:80"
      - "443:443"
    volumes:
      # Runtime routing configuration (read-only)
      - ./Runtime:/etc/terror/Runtime:ro
      # Persistent Let's Encrypt TLS certificates
      - terror_certs:/var/lib/terror/certs
      # Static site or SPA files
      - ./html:/var/www/html:ro

volumes:
  terror_certs:
```

Start the container:

```bash
docker compose up -d
```

Check container logs:

```bash
docker compose logs -f terror
```

---

### Persistent Volumes Reference

| Container Path | Type | Purpose |
| --- | --- | --- |
| `/etc/terror/Runtime` | File / Volume | Routing rules configuration file |
| `/var/lib/terror/certs` | Directory Volume | ACME TLS / Let's Encrypt certificate cache |
| `/var/www/html` | Directory Volume | Static website and SPA fallback assets |

---

### Running CLI Commands in Docker

You can run Terror Server's built-in diagnostics and validation tools directly inside the running container:

```bash
# Check route and DNS health
docker exec -it terror terror status

# Validate Runtime file syntax
docker exec -it terror terror validate

# View version info
docker exec -it terror terror version
```

### Reloading Configuration in Docker

When updating `./Runtime` on your host machine:

```bash
docker restart terror
# or with docker compose:
docker compose restart terror
```

---

### Build From Source (Dockerfile)

If you prefer building your own container image from the repository source:

```bash
git clone https://github.com/devlopersabbir/terrorserver.git
cd terrorserver
docker build -t terror:local .
```

---

## Verification

After installation, verify that Terror Server is healthy:

::: code-group
```bash [Native Systemd]
terror version
terror status
sudo systemctl status terror
```

```bash [Docker]
docker exec -it terror terror version
docker exec -it terror terror status
docker ps
```
:::

## Next Steps

1. Point your DNS `A` records to the server's public IP.
2. Ensure inbound ports `80` and `443` are open in your firewall / security groups.
3. Configure your backend applications or static sites in [Runtime File](/configuration/runtime).
4. Run `terror status` to verify end-to-end DNS, TLS, and upstream routing.
