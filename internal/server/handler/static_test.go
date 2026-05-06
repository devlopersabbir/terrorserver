package handler

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/devlopersabbir/terrorserver/internal/config"
)

func TestStaticFallback(t *testing.T) {
	// Create a temporary directory for static files
	tmpDir, err := os.MkdirTemp("", "static-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create index.html
	indexContent := "<html><body>Index</body></html>"
	err = os.WriteFile(filepath.Join(tmpDir, "index.html"), []byte(indexContent), 0644)
	if err != nil {
		t.Fatal(err)
	}

	route := config.Route{
		Root:     tmpDir,
		Fallback: "/index.html",
	}

	t.Run("existing file", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		rr := httptest.NewRecorder()

		Static(rr, req, route)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status OK, got %d", rr.Code)
		}
		if rr.Body.String() != indexContent {
			t.Errorf("expected %q, got %q", indexContent, rr.Body.String())
		}
	})

	t.Run("fallback to index.html", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/non-existent", nil)
		rr := httptest.NewRecorder()

		Static(rr, req, route)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status OK, got %d", rr.Code)
		}
		if rr.Body.String() != indexContent {
			t.Errorf("expected %q, got %q", indexContent, rr.Body.String())
		}
	})
}

func TestStaticDownload(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "static-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create install.sh
	shContent := "#!/bin/bash\necho hello"
	err = os.WriteFile(filepath.Join(tmpDir, "install.sh"), []byte(shContent), 0644)
	if err != nil {
		t.Fatal(err)
	}

	route := config.Route{
		Root: tmpDir,
	}

	t.Run("force download .sh", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/install.sh", nil)
		rr := httptest.NewRecorder()

		Static(rr, req, route)

		if rr.Code != http.StatusOK {
			t.Errorf("expected status OK, got %d", rr.Code)
		}
		
		cd := rr.Header().Get("Content-Disposition")
		expectedCD := "attachment; filename=\"install.sh\""
		if cd != expectedCD {
			t.Errorf("expected Content-Disposition %q, got %q", expectedCD, cd)
		}
	})

	t.Run("normal file .html", func(t *testing.T) {
		err = os.WriteFile(filepath.Join(tmpDir, "test.html"), []byte("test"), 0644)
		if err != nil {
			t.Fatal(err)
		}

		req := httptest.NewRequest("GET", "/test.html", nil)
		rr := httptest.NewRecorder()

		Static(rr, req, route)

		cd := rr.Header().Get("Content-Disposition")
		if cd != "" {
			t.Errorf("expected no Content-Disposition for .html, got %q", cd)
		}
	})
}
