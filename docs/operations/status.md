# Status Checks

`terror status` is the main operational command. It checks the local installation and route health.

```bash
terror status
```

Example output:

```txt
terrorserver status
-------------------------------------
ok config: /etc/terror/Runtime
ok listen: :80
ok routes: 3 configured
ok service: terror is active
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

## What It Checks

- Runtime config file.
- Configured routes.
- systemd service state.
- Expected listeners.
- DNS resolution for domains.
- TLS availability.
- Static root paths.
- Proxy upstream reachability.

## When To Run It

Run it:

- After every Runtime change.
- After installing or updating.
- Before enabling forced HTTPS redirects.
- When a domain, port, static route, or proxy route is not behaving as expected.
