# CLI Commands

The `terror` command manages serving, validation, status checks, updates, and version output.

| Command | Description |
| --- | --- |
| `terror` | Start the server |
| `terror serve` | Start the server |
| `terror start` | Start the server |
| `terror validate` | Validate Runtime config |
| `terror status` | Show service, listener, DNS, TLS, static, and upstream health |
| `terror update` | Pull and install the latest stable release |
| `terror upgrade` | Alias for `terror update` |
| `terror version` | Print binary version |
| `terror help` | Show command help |

## Short Aliases

| Alias | Command |
| --- | --- |
| `terror s` | `terror serve` |
| `terror st` | `terror status` |
| `terror v` | `terror validate` |
| `terror u` | `terror update` |

## Production Checks

Use these after Runtime edits:

```bash
terror validate
terror status
```

Use these when debugging the installed service:

```bash
sudo systemctl status terror
sudo journalctl -u terror -n 100 --no-pager
```
