#!/usr/bin/env bash
set -euo pipefail

# ─────────────────────────────────────────────
# terrorserver install.sh
# Builds and installs terrorserver as a systemd service
# ─────────────────────────────────────────────

BINARY_NAME="terror"
INSTALL_PATH="/usr/local/bin/terror"
CONFIG_DIR="/etc/terror"
CONFIG_FILE="$CONFIG_DIR/Runtime"
SERVICE_FILE="/etc/systemd/system/terror.service"
LISTEN_ADDR="${TERROR_ADDR:-:80}"

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

log_info()  { echo -e "${GREEN}[INFO]${NC} $*"; }
log_warn()  { echo -e "${YELLOW}[WARN]${NC} $*"; }
log_error() { echo -e "${RED}[ERROR]${NC} $*" >&2; }

require_root() {
  if [[ $EUID -ne 0 ]]; then
    log_error "This script must be run as root (use sudo)"
    exit 1
  fi
}

require_go() {
  if ! command -v go &>/dev/null; then
    log_error "Go is not installed. Install Go 1.21+ from https://go.dev/dl/"
    exit 1
  fi
  local goversion
  goversion=$(go version | awk '{print $3}' | sed 's/go//')
  log_info "Found Go $goversion"
}

build_binary() {
  log_info "Building $BINARY_NAME..."
  local script_dir
  script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
  local repo_root
  repo_root="$(cd "$script_dir/.." && pwd)"

  cd "$repo_root"
  go build -ldflags="-s -w" -o "$BINARY_NAME" ./cmd/terror
  log_info "Build successful"
}

install_binary() {
  log_info "Installing binary to $INSTALL_PATH"
  mv "$BINARY_NAME" "$INSTALL_PATH"
  chmod 755 "$INSTALL_PATH"
}

create_config() {
  if [[ -d "$CONFIG_DIR" ]]; then
    log_warn "Config directory $CONFIG_DIR already exists — skipping"
    return
  fi
  log_info "Creating config directory at $CONFIG_DIR"
  mkdir -p "$CONFIG_DIR"

  log_info "Writing example config to $CONFIG_FILE"
  cat > "$CONFIG_FILE" <<'EOF'
# terrorserver Runtime config
# Edit this file — changes are reloaded automatically (no restart needed)

# example.com {
#     proxy localhost:3000
# }

# static.example.com {
#     root /var/www/html
#     file_server
# }

# :8080 {
#     proxy localhost:8081
# }
EOF
  chmod 644 "$CONFIG_FILE"
}

install_service() {
  log_info "Creating systemd service at $SERVICE_FILE"
  cat > "$SERVICE_FILE" <<EOF
[Unit]
Description=terrorserver — minimal domain router & reverse proxy
Documentation=https://github.com/terrorserver/terror
After=network.target
Wants=network-online.target

[Service]
Type=simple
ExecStart=$INSTALL_PATH serve
Restart=always
RestartSec=3
Environment=TERROR_CONFIG=$CONFIG_FILE
Environment=TERROR_ADDR=$LISTEN_ADDR

# Security hardening
NoNewPrivileges=yes
PrivateTmp=yes
ProtectSystem=strict
ReadWritePaths=$CONFIG_DIR
CapabilityBoundingSet=CAP_NET_BIND_SERVICE
AmbientCapabilities=CAP_NET_BIND_SERVICE

[Install]
WantedBy=multi-user.target
EOF

  systemctl daemon-reload
  systemctl enable terror
  systemctl start terror
  log_info "Service started. Check status with: systemctl status terror"
}

print_success() {
  echo ""
  echo -e "${GREEN}────────────────────────────────────────${NC}"
  echo -e "${GREEN}  terrorserver installed successfully!${NC}"
  echo -e "${GREEN}────────────────────────────────────────${NC}"
  echo ""
  echo "  Config file:    $CONFIG_FILE"
  echo "  Listen address: $LISTEN_ADDR"
  echo "  Binary:         $INSTALL_PATH"
  echo ""
  echo "  Useful commands:"
  echo "    terror validate          — check config syntax"
  echo "    terror status            — show status"
  echo "    systemctl status terror  — systemd status"
  echo "    journalctl -u terror -f  — live logs"
  echo ""
  echo "  Edit the config file and changes apply instantly (no restart)."
  echo ""
}

main() {
  require_root
  require_go
  build_binary
  install_binary
  create_config
  install_service
  print_success
}

main "$@"
