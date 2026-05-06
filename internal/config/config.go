package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// RouteType defines the handler type for a route.
type RouteType string

const (
	RouteProxy  RouteType = "proxy"
	RouteStatic RouteType = "static"
)

// Route represents a single parsed routing entry.
type Route struct {
	Host     string    // "api.example.com" or ":4000"
	Type     RouteType // proxy | static
	Target   string    // upstream address for proxy
	Root     string    // filesystem root for static
	Fallback string    // fallback file for static (e.g. index.html)
}

// Config holds all parsed routes.
type Config struct {
	Routes []Route
}

// Parse reads and parses the terrorserver Runtime config file.
func Parse(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("cannot open config file %q: %w", path, err)
	}
	defer f.Close()

	cfg := &Config{}
	scanner := bufio.NewScanner(f)

	var (
		inBlock    bool
		currentHost string
		hasProxy   bool
		hasStatic  bool
		proxyTarget    string
		staticRoot     string
		staticFallback string
		lineNum        int
	)

	flushBlock := func() error {
		if currentHost == "" {
			return nil
		}
		if hasProxy && hasStatic {
			return fmt.Errorf("block %q has both proxy and file_server — only one allowed", currentHost)
		}
		if !hasProxy && !hasStatic {
			return fmt.Errorf("block %q has no proxy or file_server directive", currentHost)
		}
		r := Route{Host: currentHost}
		if hasProxy {
			r.Type = RouteProxy
			r.Target = proxyTarget
		} else {
			r.Type = RouteStatic
			r.Root = staticRoot
			r.Fallback = staticFallback
		}
		cfg.Routes = append(cfg.Routes, r)
		return nil
	}

	resetBlock := func() {
		inBlock = false
		currentHost = ""
		hasProxy = false
		hasStatic = false
		proxyTarget = ""
		staticRoot = ""
		staticFallback = ""
	}

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if !inBlock {
			// Expect: "host.name {" or ":port {"
			if strings.HasSuffix(line, "{") {
				host := strings.TrimSpace(strings.TrimSuffix(line, "{"))
				host = strings.TrimSpace(host)
				if host == "" {
					return nil, fmt.Errorf("line %d: empty block host", lineNum)
				}
				currentHost = host
				inBlock = true
				continue
			}
			return nil, fmt.Errorf("line %d: unexpected token outside block: %q", lineNum, line)
		}

		// Inside block
		if line == "}" {
			if err := flushBlock(); err != nil {
				return nil, fmt.Errorf("line %d: %w", lineNum, err)
			}
			resetBlock()
			continue
		}

		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		directive := strings.ToLower(parts[0])
		switch directive {
		case "proxy":
			if len(parts) < 2 {
				return nil, fmt.Errorf("line %d: proxy directive requires an upstream address", lineNum)
			}
			proxyTarget = parts[1]
			hasProxy = true

		case "root":
			if len(parts) < 2 {
				return nil, fmt.Errorf("line %d: root directive requires a path", lineNum)
			}
			staticRoot = parts[1]

		case "file_server":
			hasStatic = true
			if len(parts) >= 3 && parts[1] == "{path}" {
				staticFallback = parts[2]
			}

		default:
			return nil, fmt.Errorf("line %d: unknown directive %q", lineNum, directive)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	if inBlock {
		return nil, fmt.Errorf("unclosed block for host %q", currentHost)
	}

	return cfg, nil
}

// RouteMap builds a lookup map from the config.
// Key is normalized host (lowercase).
func (c *Config) RouteMap() map[string]Route {
	m := make(map[string]Route, len(c.Routes))
	for _, r := range c.Routes {
		m[strings.ToLower(r.Host)] = r
	}
	return m
}
