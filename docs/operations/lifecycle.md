# Updates And Uninstall

Terror Server uses the hosted installer for updates and a hosted uninstaller for removal.

## Update

Update an installed server to the newest stable release:

```bash
terror update
```

`terror upgrade` is an alias:

```bash
terror upgrade
```

The update command runs:

```txt
https://terror.softvenceomega.com/install.sh
```

The installer downloads the latest GitHub release asset, replaces `/usr/local/bin/terror`, refreshes the systemd service files, and keeps existing Runtime config and welcome files when they already exist.

## Uninstall

Interactive uninstall:

```bash
curl -fsSL https://terror.softvenceomega.com/uninstall.sh | sudo bash
```

Non-interactive uninstall:

```bash
curl -fsSL https://terror.softvenceomega.com/uninstall.sh | sudo bash -s -- --yes
```

The uninstaller stops and disables the systemd service, removes the binary and watcher units, then asks before deleting config, welcome site, and certificate cache directories.
