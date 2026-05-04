# Troubleshooting

Start with:

```bash
terror validate
terror status
sudo journalctl -u terror -n 120 --no-pager
```

## Domain Does Not Route

Check DNS and local host routing:

```bash
nslookup app.example.com
curl -H "Host: app.example.com" http://127.0.0.1
terror status
```

If the local curl works but the browser does not, the issue is usually DNS propagation, firewall rules, cloud security group rules, or another service in front of the server.

## HTTPS Shows As Not Secure

Check the TLS listener and service logs:

```bash
terror status
sudo journalctl -u terror -n 120 --no-pager
curl -vk https://app.example.com
```

Let's Encrypt needs public access to port `80` for the HTTP challenge and port `443` for HTTPS traffic.

## New Port Route Does Not Work

Confirm the Runtime file was saved and the path watcher is active:

```bash
sudo systemctl status terror.path
sudo journalctl -u terror -n 80 --no-pager
terror status
```

Adding or removing listener ports requires a process restart because ports must be rebound. The installer configures this restart automatically through `terror.path`.

## Upstream Is Unreachable

Check the app directly from the server:

```bash
curl http://127.0.0.1:4000
terror status
```

Use `localhost:4000` or `127.0.0.1:4000` in Runtime when the app is running on the same machine.

## Static Route Returns The Wrong Page

Check the configured `root` and confirm the site build exists:

```bash
ls -la /var/www/app
terror status
```

Remember that missing static paths fall back to `/`, which is useful for client-side routing but can hide a missing asset path during debugging.
