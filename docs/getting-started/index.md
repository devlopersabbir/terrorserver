# Getting Started

Terror Server runs as a single Go binary and is usually installed as a systemd-managed service. The normal production flow is:

1. Install the latest stable release.
2. Edit `/etc/terror/Runtime`.
3. Validate the config.
4. Let the systemd watcher restart `terror.service`.
5. Check route health with `terror status`.

## Install

```bash
curl -fsSL https://terror.softvenceomega.com/install.sh | sudo bash
```

Then check the service:

```bash
terror status
```

::: tip
The installer enables `terror.path`, a systemd path watcher. Saving `/etc/terror/Runtime` automatically triggers a service restart so new domains, removed routes, ports, and TLS listeners are picked up.
:::

## Your First Runtime

Open the Runtime file:

```bash
sudo vim /etc/terror/Runtime
```

Start with a default static site and one proxied app:

```txt
:80 {
    root /var/www/terrorserver
    file_server
}

app.example.com {
    proxy localhost:4000
}
```

Validate before saving a production change:

```bash
terror validate
```

After saving, confirm the route:

```bash
terror status
```

## Mental Model

Each Runtime block has an address and exactly one handler.

| Address type | Example | What it matches |
| --- | --- | --- |
| Domain | `app.example.com` | Requests with that hostname |
| Port | `:9090` | Any request received on that listener |

Each block chooses either `proxy` or `file_server`.

```txt
api.example.com {
    proxy localhost:3000
}

docs.example.com {
    root /var/www/docs
    file_server
}
```

## Next Steps

- Use [Runtime File](/configuration/runtime) for all syntax rules.
- Use [Reverse Proxy](/guides/reverse-proxy) to publish local apps.
- Use [Static Sites](/guides/static-sites) for frontend builds and simple websites.
- Use [HTTPS and Domains](/guides/https-domains) before turning on forced redirects.
