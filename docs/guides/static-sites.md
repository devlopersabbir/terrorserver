# Static Sites

Use `root` plus `file_server` to serve a directory.

```txt
example.com {
    root /var/www/example
    file_server
}
```

## Serve A Frontend Build

Build your frontend, copy the output to a server directory, then point Terror Server at it.

```txt
app.example.com {
    root /var/www/app
    file_server
}
```

The static handler falls back to `/` when a requested file does not exist. That makes single-page apps with client-side routing work without extra rewrite rules.

## Default Welcome Site

The installer creates the default welcome root:

```txt
/var/www/terrorserver
```

A simple default Runtime can serve it on port `80`:

```txt
:80 {
    root /var/www/terrorserver
    file_server
}
```

## Validate The Root

Before routing traffic, confirm the directory exists and contains an `index.html`:

```bash
ls -la /var/www/app
terror validate
terror status
```

If `terror status` reports a static root problem, fix the path or permissions first.
