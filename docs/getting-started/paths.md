# Default Paths

Terror Server keeps its production footprint small and predictable.

| Path | Purpose |
| --- | --- |
| `/usr/local/bin/terror` | Installed binary |
| `/etc/terror/Runtime` | Main Runtime configuration |
| `/var/www/terrorserver` | Default static welcome site root |
| `/var/www/terrorserver/index.html` | Default welcome page |
| `/var/lib/terror/certs` | Let's Encrypt certificate cache |
| `/etc/systemd/system/terror.service` | Main systemd service |
| `/etc/systemd/system/terror.path` | Runtime change watcher |
| `/etc/systemd/system/terror-restart.service` | Restart helper triggered by `terror.path` |

## Config Path

The service reads `/etc/terror/Runtime` by default. Override it with:

```bash
TERROR_CONFIG=/path/to/Runtime terror serve
```

For installed service overrides, use:

```bash
sudo systemctl edit terror
```

Then add:

```ini
[Service]
Environment=TERROR_CONFIG=/etc/terror/Runtime
```

Reload systemd after changing service overrides:

```bash
sudo systemctl daemon-reload
sudo systemctl restart terror
```
