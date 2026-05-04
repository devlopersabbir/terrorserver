# Runtime Examples

This page collects copyable Runtime patterns.

## Default Static Site

```txt
:80 {
    root /var/www/terrorserver
    file_server
}
```

## Domain To Local App

```txt
api.example.com {
    proxy localhost:3000
}
```

## Port To Local App

```txt
:9090 {
    proxy localhost:4000
}
```

## Static Domain

```txt
static.example.com {
    root /var/www/html
    file_server
}
```

## Multi-App Server

```txt
app.example.com {
    proxy localhost:4000
}

api.example.com {
    proxy localhost:3000
}

docs.example.com {
    root /var/www/docs
    file_server
}

:9090 {
    proxy localhost:9000
}
```

## Validate

```bash
terror validate
terror status
```
