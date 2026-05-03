#!/usr/bin/env bash
set -euo pipefail

# ─────────────────────────────────────────────
# terrorserver install.sh
# Installs terrorserver from the stable GitHub release as a systemd service
# ─────────────────────────────────────────────

BINARY_NAME="terror"
INSTALL_PATH="/usr/local/bin/terror"
CONFIG_DIR="/etc/terror"
CONFIG_FILE="$CONFIG_DIR/Runtime"
WEB_ROOT="${TERROR_WEB_ROOT:-/var/www/terrorserver}"
WEB_INDEX="$WEB_ROOT/index.html"
SERVICE_FILE="/etc/systemd/system/terror.service"
LISTEN_ADDR="${TERROR_ADDR:-:80}"
REPO="${TERROR_REPO:-devlopersabbir/terrorserver}"
CHANNEL="${TERROR_CHANNEL:-stable}"
DOWNLOAD_BASE="https://github.com/$REPO/releases/latest/download"
AUTHOR_NAME="Sabbir Hossain Shuvo"
AUTHOR_URL="https://devlopersabbir.github.io"

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

require_downloader() {
  if command -v curl &>/dev/null; then
    DOWNLOADER="curl"
    return
  fi
  if command -v wget &>/dev/null; then
    DOWNLOADER="wget"
    return
  fi
  log_error "curl or wget is required to download the GitHub release"
  exit 1
}

detect_asset() {
  local os arch
  os="$(uname -s | tr '[:upper:]' '[:lower:]')"
  arch="$(uname -m)"

  if [[ "$os" != "linux" ]]; then
    log_error "This installer supports Linux systemd hosts only (detected: $os)"
    exit 1
  fi

  case "$arch" in
    x86_64|amd64) arch="amd64" ;;
    aarch64|arm64) arch="arm64" ;;
    *)
      log_error "Unsupported CPU architecture: $arch"
      exit 1
      ;;
  esac

  RELEASE_ASSET="${TERROR_ASSET:-$BINARY_NAME-$os-$arch}"
}

download_binary() {
  local url tmp_file
  url="$DOWNLOAD_BASE/$RELEASE_ASSET"
  tmp_file="$(mktemp)"

  log_info "Pulling (-$CHANNEL) release from github"
  log_info "Downloading $url"

  if [[ "$DOWNLOADER" == "curl" ]]; then
    curl -fsSL "$url" -o "$tmp_file"
  else
    wget -qO "$tmp_file" "$url"
  fi

  chmod 755 "$tmp_file"
  DOWNLOADED_BINARY="$tmp_file"
}

install_binary() {
  log_info "Installing binary to $INSTALL_PATH"
  mv "$DOWNLOADED_BINARY" "$INSTALL_PATH"
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
  cat > "$CONFIG_FILE" <<EOF
# terrorserver Runtime config
# Edit this file — changes are reloaded automatically (no restart needed)

:80 {
    root $WEB_ROOT
    file_server
}
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

create_welcome_site() {
  log_info "Creating default welcome page at $WEB_INDEX"
  mkdir -p "$WEB_ROOT"

  if [[ -f "$WEB_INDEX" ]]; then
    log_warn "Welcome page $WEB_INDEX already exists — skipping"
    return
  fi

  cat > "$WEB_INDEX" <<'EOF'
<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Terror Server</title>
  <style>
    :root {
      color-scheme: dark;
      --bg: #101113;
      --panel: #181b1f;
      --text: #f4f5f7;
      --muted: #a5adba;
      --line: #2b3038;
      --accent: #ef4444;
      --accent-soft: rgba(239, 68, 68, .16);
    }
    * { box-sizing: border-box; }
    html, body { height: 100%; }
    body {
      margin: 0;
      display: grid;
      place-items: center;
      min-height: 100%;
      background: var(--bg);
      color: var(--text);
      font-family: Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
    }
    main {
      width: min(680px, calc(100% - 32px));
      padding: 40px;
      border: 1px solid var(--line);
      border-radius: 8px;
      background: var(--panel);
      box-shadow: 0 24px 80px rgba(0, 0, 0, .34);
    }
    .mark {
      display: inline-flex;
      align-items: center;
      justify-content: center;
      width: 44px;
      height: 44px;
      margin-bottom: 22px;
      border-radius: 8px;
      background: var(--accent-soft);
      color: var(--accent);
      font-size: 24px;
      font-weight: 800;
    }
    h1 {
      margin: 0 0 12px;
      font-size: 40px;
      line-height: 1.1;
      font-weight: 800;
      letter-spacing: 0;
    }
    p {
      margin: 0;
      color: var(--muted);
      font-size: 17px;
      line-height: 1.65;
    }
    code {
      display: inline-block;
      margin-top: 26px;
      padding: 10px 12px;
      border: 1px solid var(--line);
      border-radius: 6px;
      background: #0c0d0f;
      color: #ffffff;
      font-size: 14px;
    }
  </style>
</head>
<body>
  <main>
    <div class="mark">T</div>
    <h1>Terror Server is running</h1>
    <p>Your server is online. Add a route in your Runtime config to serve your own site.</p>
    <code>/etc/terror/Runtime</code>
  </main>
</body>
</html>
EOF

  chmod 755 "$WEB_ROOT"
  chmod 644 "$WEB_INDEX"
}

install_service() {
  log_info "Creating systemd service at $SERVICE_FILE"
  cat > "$SERVICE_FILE" <<EOF
[Unit]
Description=terrorserver — minimal domain router & reverse proxy
Documentation=https://github.com/devlopersabbir/terrorserver
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
  echo "  Welcome page:   $WEB_INDEX"
  echo ""
  echo "  Useful commands:"
  echo "    terror validate          — check config syntax"
  echo "    terror status            — show status"
  echo "    systemctl status terror  — systemd status"
  echo "    journalctl -u terror -f  — live logs"
  echo ""
  echo "  Edit the config file and changes apply instantly (no restart)."
  echo ""
  echo "  Built by: $AUTHOR_NAME"
  echo "  Portfolio:   $AUTHOR_URL"
  echo "  Project:  https://github.com/$REPO"
  echo ""
  echo "  If you find any issues or have suggestions, feel free to raise a pull request."
}

main() {
  require_root
  require_downloader
  detect_asset
  download_binary
  install_binary
  create_welcome_site
  create_config
  install_service
  print_success
}

main "$@"
