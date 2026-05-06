package main

import (
	"os"
	"testing"
)

func TestConfigPathDefault(t *testing.T) {
	os.Unsetenv("TERROR_CONFIG")
	if got := configPath(); got != defaultConfigPath {
		t.Errorf("configPath() = %q, want %q", got, defaultConfigPath)
	}
}

func TestConfigPathEnvOverride(t *testing.T) {
	const want = "/tmp/terror-config"
	t.Setenv("TERROR_CONFIG", want)
	if got := configPath(); got != want {
		t.Errorf("configPath() = %q, want %q", got, want)
	}
}

func TestListenAddrDefault(t *testing.T) {
	os.Unsetenv("TERROR_ADDR")
	if got := listenAddr(); got != defaultAddr {
		t.Errorf("listenAddr() = %q, want %q", got, defaultAddr)
	}
}

func TestListenAddrEnvOverride(t *testing.T) {
	const want = ":9090"
	t.Setenv("TERROR_ADDR", want)
	if got := listenAddr(); got != want {
		t.Errorf("listenAddr() = %q, want %q", got, want)
	}
}
