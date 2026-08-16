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

```bash
curl -fsSL https://raw.githubusercontent.com/devlopersabbir/terrorserver/main/scripts/install.sh | sudo bash
```

The installer re-fetches the binary and service files, leaving existing config and root files intact if present.

## Uninstalling

Interactive removal:

```bash
curl -fsSL https://raw.githubusercontent.com/devlopersabbir/terrorserver/main/scripts/uninstall.sh | sudo bash
```

Unattended removal:

```bash
curl -fsSL https://raw.githubusercontent.com/devlopersabbir/terrorserver/main/scripts/uninstall.sh | sudo bash -s -- --yes
```

The uninstaller stops and disables the systemd service, removes the binary and watcher units, then asks before deleting config, welcome site, and certificate cache directories.
