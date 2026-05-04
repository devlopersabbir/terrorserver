# Install

The hosted installer downloads the latest stable Linux release, installs the binary, creates service files, and prepares the default Runtime and welcome site paths.

## Standard Install

```bash
curl -fsSL https://terror.softvenceomega.com/install.sh | sudo bash
```

The installer configures:

- `/usr/local/bin/terror`
- `/etc/terror/Runtime`
- `/var/www/terrorserver`
- `terror.service`
- `terror.path`
- `terror-restart.service`

## Verify The Install

```bash
terror version
terror status
sudo systemctl status terror
sudo systemctl status terror.path
```

The service should be active, and the path watcher should be enabled so Runtime edits restart the server.

## Installer Inputs

The installer is controlled by environment variables when you need a non-default release source or template host.

| Variable | Default | Purpose |
| --- | --- | --- |
| `TERROR_REPO` | `devlopersabbir/terrorserver` | GitHub repository used for release downloads |
| `TERROR_ASSET` | auto-detected | Release asset override |
| `TERROR_CHANNEL` | `stable` | Display channel label |
| `TERROR_TEMPLATE_BASE` | `https://terror.softvenceomega.com` | Host for `Runtime` and `welcome.html` templates |
| `TERROR_WEB_ROOT` | `/var/www/terrorserver` | Default welcome page root |

## After Install

1. Point DNS records at the server.
2. Open inbound ports `80` and `443` when using domain routes.
3. Put local apps behind `localhost` ports.
4. Edit `/etc/terror/Runtime`.
5. Run `terror validate` and `terror status`.
