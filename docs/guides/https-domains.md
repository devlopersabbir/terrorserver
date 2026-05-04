# HTTPS And Domains

For domain routes, Terror Server enables automatic Let's Encrypt TLS by default.

```txt
app.example.com {
    proxy localhost:4000
}
```

When TLS is enabled:

- Port `443` serves HTTPS for domain routes.
- Port `80` handles ACME HTTP challenges.
- Normal HTTP domain traffic stays available by default while HTTPS is served on port `443`.
- HTTP to HTTPS redirects can be enabled after certificate issuance is confirmed.
- IP-based and port-based routes such as `:9090` stay on HTTP.

## Requirements

- The domain must resolve to the server's public IP.
- Inbound ports `80` and `443` must be open.
- No other service can already bind ports `80` or `443`.
- The installed service must have write access to `/var/lib/terror/certs`.

## Disable Automatic TLS

Only disable TLS when another layer manages certificates or you are running in an environment that cannot satisfy Let's Encrypt challenges.

```bash
sudo systemctl edit terror
```

Add:

```ini
[Service]
Environment=TERROR_AUTO_TLS=false
```

Reload:

```bash
sudo systemctl daemon-reload
sudo systemctl restart terror
```

## Force HTTPS Redirects

Enable redirects only after `https://your-domain` is confirmed healthy.

```ini
[Service]
Environment=TERROR_HTTPS_REDIRECT=true
```

Then reload and restart:

```bash
sudo systemctl daemon-reload
sudo systemctl restart terror
```

## Check TLS Health

```bash
terror status
curl -vk https://app.example.com
sudo journalctl -u terror -n 120 --no-pager
```
