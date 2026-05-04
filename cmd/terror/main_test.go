package main

import "testing"

func TestInstallURLDefault(t *testing.T) {
	t.Setenv("TERROR_INSTALL_URL", "")

	if got := installURL(); got != defaultInstallURL {
		t.Fatalf("installURL() = %q, want %q", got, defaultInstallURL)
	}
}

func TestInstallURLEnvOverride(t *testing.T) {
	const want = "https://example.com/install.sh"
	t.Setenv("TERROR_INSTALL_URL", want)

	if got := installURL(); got != want {
		t.Fatalf("installURL() = %q, want %q", got, want)
	}
}
