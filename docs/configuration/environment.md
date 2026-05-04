# Environment Variables

Use environment variables for service paths, TLS behavior, installer behavior, and update sources.

| Variable | Default | Used by | Description |
| --- | --- | --- | --- |
| `TERROR_CONFIG` | `/etc/terror/Runtime` | Binary, service | Runtime config path |
| `TERROR_ADDR` | `:80` | Binary, installer | Default listen address |
| `TERROR_CERT_CACHE` | `/var/lib/terror/certs` | Binary, installer | Let's Encrypt certificate cache |
| `TERROR_AUTO_TLS` | enabled | Binary | Set to `false`, `0`, or `no` to disable automatic TLS |
| `TERROR_HTTPS_REDIRECT` | disabled | Binary | Set to `true`, `1`, or `yes` to redirect domain HTTP traffic to HTTPS |
| `TERROR_INSTALL_URL` | `https://terror.softvenceomega.com/install.sh` | `terror update` | Installer URL used by update command |
| `TERROR_WEB_ROOT` | `/var/www/terrorserver` | Installer, uninstaller | Default welcome page root |
| `TERROR_TEMPLATE_BASE` | `https://terror.softvenceomega.com` | Installer | Host for `Runtime` and `welcome.html` templates |
| `TERROR_REPO` | `devlopersabbir/terrorserver` | Installer | GitHub repository used for release downloads |
| `TERROR_ASSET` | auto-detected | Installer | Release asset override |
| `TERROR_CHANNEL` | `stable` | Installer | Display channel label |

## Set Service Overrides

Use systemd drop-ins for installed services:

```bash
sudo systemctl edit terror
```

Example:

```ini
[Service]
Environment=TERROR_AUTO_TLS=false
Environment=TERROR_HTTPS_REDIRECT=false
```

Apply the change:

```bash
sudo systemctl daemon-reload
sudo systemctl restart terror
```

## TLS Redirects

Only enable redirects after `https://your-domain` is confirmed healthy:

```ini
[Service]
Environment=TERROR_HTTPS_REDIRECT=true
```

This avoids redirecting traffic to HTTPS before a certificate and listener are working.
