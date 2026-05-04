# Systemd

The installer configures Terror Server as a systemd service with a Runtime file watcher.

## Common Commands

```bash
sudo systemctl status terror
sudo journalctl -u terror -n 100 --no-pager
sudo journalctl -uf terror
sudo systemctl restart terror
sudo systemctl status terror.path
```

## Runtime Watcher

`terror.path` watches `/etc/terror/Runtime`. When the file changes, it triggers `terror-restart.service`, which restarts `terror.service`.

This matters because adding or removing listener ports requires the process to restart and bind the new socket set.

## Service Hardening

The installed service uses:

- `Restart=always`
- `CAP_NET_BIND_SERVICE` for binding low ports.
- `ProtectSystem=strict`
- `ReadWritePaths=/etc/terror /var/lib/terror`
- `PrivateTmp=yes`
- `NoNewPrivileges=yes`

## Editing Service Environment

```bash
sudo systemctl edit terror
```

Example:

```ini
[Service]
Environment=TERROR_AUTO_TLS=false
```

Apply:

```bash
sudo systemctl daemon-reload
sudo systemctl restart terror
```
