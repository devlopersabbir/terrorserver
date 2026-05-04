# Runtime File

The Runtime file is the routing source of truth. It uses compact blocks where the block label is either a domain name or a port listener.

```txt
:80 {
    root /var/www/terrorserver
    file_server
}

app.example.com {
    proxy localhost:4000
}
```

## Block Shape

```txt
address {
    directive value
    handler
}
```

| Part | Meaning |
| --- | --- |
| `address` | A domain such as `app.example.com` or a port such as `:9090` |
| `proxy` | Reverse proxy handler with an upstream address |
| `root` | Static file root used by `file_server` |
| `file_server` | Static file handler |

## Rules

- One block per domain or port.
- A block can use `proxy` or `file_server`, not both.
- `file_server` requires a `root` directive.
- `proxy` requires an upstream address.
- Comments start with `#`.
- Unknown directives fail validation.
- Domains are matched case-insensitively.

Run validation before production edits:

```bash
terror validate
```

## Domain Routes

Domain routes match by hostname.

```txt
api.example.com {
    proxy localhost:3000
}
```

For domain routes, automatic Let's Encrypt TLS is enabled by default unless disabled with `TERROR_AUTO_TLS=false`.

## Port Routes

Port routes match requests received on that listener.

```txt
:9090 {
    proxy localhost:4000
}
```

Port routes stay on HTTP. They are useful for internal tools, private listeners, or apps that are exposed through another upstream layer.

## Static Routes

```txt
static.example.com {
    root /var/www/html
    file_server
}
```

When a requested static path does not exist, Terror Server falls back to `/`. This keeps frontend apps with client-side routing working cleanly.

## Complete Example

```txt
# Default welcome site
:80 {
    root /var/www/terrorserver
    file_server
}

# Public app over automatic HTTPS
app.example.com {
    proxy localhost:4000
}

# Internal HTTP listener
:9090 {
    proxy localhost:4000
}

# Static site
static.example.com {
    root /var/www/html
    file_server
}
```
