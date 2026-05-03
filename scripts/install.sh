#!/usr/bin/env bash
set -euo pipefail

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
TEMPLATE_BASE="${TERROR_TEMPLATE_BASE:-https://terror.softvenceomega.com}"
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

download_file() {
  local url dest
  url="$1"
  dest="$2"

  if [[ "$DOWNLOADER" == "curl" ]]; then
    curl -fsSL "$url" -o "$dest"
  else
    wget -qO "$dest" "$url"
  fi
}

escape_sed() {
  printf '%s' "$1" | sed 's/[\/&]/\\&/g'
}

install_binary() {
  log_info "Installing binary to $INSTALL_PATH"
  mv "$DOWNLOADED_BINARY" "$INSTALL_PATH"
  chmod 755 "$INSTALL_PATH"
}

create_config() {
  if [[ -f "$CONFIG_FILE" ]]; then
    log_warn "Config file $CONFIG_FILE already exists — skipping"
    return
  fi

  log_info "Creating config directory at $CONFIG_DIR"
  mkdir -p "$CONFIG_DIR"

  local template
  template="$(mktemp)"

  log_info "Pulling Runtime template from $TEMPLATE_BASE/Runtime"
  download_file "$TEMPLATE_BASE/Runtime" "$template"

  log_info "Writing example config to $CONFIG_FILE"
  sed \
    -e "s/{{LISTEN_ADDR}}/$(escape_sed "$LISTEN_ADDR")/g" \
    -e "s/{{WEB_ROOT}}/$(escape_sed "$WEB_ROOT")/g" \
    -e "s/{{AUTHOR_NAME}}/$(escape_sed "$AUTHOR_NAME")/g" \
    -e "s/{{AUTHOR_URL}}/$(escape_sed "$AUTHOR_URL")/g" \
    "$template" > "$CONFIG_FILE"

  rm -f "$template"
  chmod 644 "$CONFIG_FILE"
}

create_welcome_site() {
  log_info "Creating default welcome page at $WEB_INDEX"
  mkdir -p "$WEB_ROOT"

  if [[ -f "$WEB_INDEX" ]]; then
    log_warn "Welcome page $WEB_INDEX already exists — skipping"
    return
  fi

  log_info "Pulling welcome page from $TEMPLATE_BASE/welcome.html"
  download_file "$TEMPLATE_BASE/welcome.html" "$WEB_INDEX"

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
  local CYAN='\033[0;36m'
  local GRAY='\033[0;90m'
  local WHITE='\033[1;37m'
  local BOLD='\033[1m'

  echo -e ""
  echo -e "  ${RED}${BOLD}▲ TERRORSERVER${NC} ${GRAY}Installed${NC}"
  echo -e "  ${GREEN}Successfully installed to your system${NC}"
  echo -e ""
  
  echo -e "  ${WHITE}${BOLD}RESOURCES${NC}"
  echo -e "  ${GRAY}├─${NC} ${BOLD}Endpoint:${NC}  ${CYAN}http://localhost${LISTEN_ADDR}${NC}"
  echo -e "  ${GRAY}├─${NC} ${BOLD}Config:${NC}    ${GRAY}${CONFIG_FILE}${NC}"
  echo -e "  ${GRAY}└─${NC} ${BOLD}Web Root:${NC}  ${GRAY}${WEB_ROOT}${NC}"
  echo -e ""

  echo -e "  ${WHITE}${BOLD}CONTROL${NC}"
  echo -e "  ${GRAY}•${NC} ${CYAN}terror validate${NC}    ${GRAY}Verify config syntax${NC}"
  echo -e "  ${GRAY}•${NC} ${CYAN}terror status${NC}      ${GRAY}Check proxy health${NC}"
  echo -e "  ${GRAY}•${NC} ${CYAN}journalctl -uf terror${NC}"
  echo -e ""

  echo -e "  ${WHITE}${BOLD}MAINTAINER${NC}"
  echo -e "  ${GRAY}By${NC} ${AUTHOR_NAME} ${GRAY}(${AUTHOR_URL})${NC}"
  echo -e "  ${GRAY}Docs: https://terror.softvenceomega.com/docs${NC}"
  echo -e ""
  echo -e "  ${RED}${BOLD}watch:${NC} ${GRAY}Hot-reload is active. Keep an eye on your traffic.${NC}"
  echo -e ""
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
