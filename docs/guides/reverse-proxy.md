# Reverse Proxy

Use `proxy` when an app already listens on the server, usually on `localhost`.

```txt
api.example.com {
    proxy localhost:3000
}
```

## Domain Proxy

1. Run your app on a local port.
2. Point the domain DNS record at the server.
3. Add a domain block.
4. Validate and save the Runtime file.
5. Check `terror status`.

```txt
app.example.com {
    proxy localhost:4000
}
```

```bash
terror validate
terror status
```

## Port Proxy

Use a port block when the listener itself is the route.

```txt
:9090 {
    proxy localhost:4000
}
```

This exposes the upstream on `http://server-ip:9090`.

## Upstream Checks

If a proxy route is unhealthy, check the upstream directly from the same server:

```bash
curl http://127.0.0.1:4000
terror status
```

Use `localhost:4000` or `127.0.0.1:4000` when the app is running on the same machine.

## Common App Pairings

| App type | Runtime upstream |
| --- | --- |
| Node or Bun app on port 3000 | `proxy localhost:3000` |
| Go service on port 8080 | `proxy localhost:8080` |
| Python app on port 8000 | `proxy localhost:8000` |
| Internal dashboard on port 4000 | `:9090 { proxy localhost:4000 }` |
