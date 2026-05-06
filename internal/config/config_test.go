package config

import (
	"os"
	"testing"
)

func writeTemp(t *testing.T, content string) string {
	t.Helper()
	f, err := os.CreateTemp("", "terror-config-*.txt")
	if err != nil {
		t.Fatal(err)
	}
	f.WriteString(content)
	f.Close()
	return f.Name()
}

func TestParseProxy(t *testing.T) {
	cfg := `
api.example.com {
    proxy localhost:5000
}
`
	path := writeTemp(t, cfg)
	defer os.Remove(path)

	c, err := Parse(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(c.Routes) != 1 {
		t.Fatalf("expected 1 route, got %d", len(c.Routes))
	}
	r := c.Routes[0]
	if r.Host != "api.example.com" {
		t.Errorf("wrong host: %s", r.Host)
	}
	if r.Type != RouteProxy {
		t.Errorf("expected proxy type")
	}
	if r.Target != "localhost:5000" {
		t.Errorf("wrong target: %s", r.Target)
	}
}

func TestParseStatic(t *testing.T) {
	cfg := `
app.example.com {
    root /var/www/html
    file_server
}
`
	path := writeTemp(t, cfg)
	defer os.Remove(path)

	c, err := Parse(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	r := c.Routes[0]
	if r.Type != RouteStatic {
		t.Errorf("expected static type")
	}
	if r.Root != "/var/www/html" {
		t.Errorf("wrong root: %s", r.Root)
	}
	if r.Fallback != "" {
		t.Errorf("expected empty fallback, got %s", r.Fallback)
	}
}

func TestParseStaticWithFallback(t *testing.T) {
	cfg := `
spa.example.com {
    root /var/www/spa
    file_server {path} /index.html
}
`
	path := writeTemp(t, cfg)
	defer os.Remove(path)

	c, err := Parse(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	r := c.Routes[0]
	if r.Type != RouteStatic {
		t.Errorf("expected static type")
	}
	if r.Fallback != "/index.html" {
		t.Errorf("expected fallback /index.html, got %s", r.Fallback)
	}
}

func TestParsePortBlock(t *testing.T) {
	cfg := `
:4000 {
    proxy localhost:3000
}
`
	path := writeTemp(t, cfg)
	defer os.Remove(path)

	c, err := Parse(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.Routes[0].Host != ":4000" {
		t.Errorf("expected :4000, got %s", c.Routes[0].Host)
	}
}

func TestParseBothDirectivesError(t *testing.T) {
	cfg := `
bad.example.com {
    proxy localhost:5000
    root /var/www
    file_server
}
`
	path := writeTemp(t, cfg)
	defer os.Remove(path)

	_, err := Parse(path)
	if err == nil {
		t.Fatal("expected error for both proxy and file_server in same block")
	}
}

func TestParseMultipleBlocks(t *testing.T) {
	cfg := `
a.example.com {
    proxy localhost:3001
}

b.example.com {
    root /srv/b
    file_server
}

:9090 {
    proxy localhost:9091
}
`
	path := writeTemp(t, cfg)
	defer os.Remove(path)

	c, err := Parse(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(c.Routes) != 3 {
		t.Fatalf("expected 3 routes, got %d", len(c.Routes))
	}
}
