#!/usr/bin/env bash
set -euo pipefail

# ─────────────────────────────────────────────
# terrorserver uninstall.sh
# ─────────────────────────────────────────────

INSTALL_PATH="/usr/local/bin/terror"
CONFIG_DIR="/etc/terror"
SERVICE_FILE="/etc/systemd/system/terror.service"

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

log_info()  { echo -e "${GREEN}[INFO]${NC} $*"; }
log_warn()  { echo -e "${YELLOW}[WARN]${NC} $*"; }

require_root() {
  if [[ $EUID -ne 0 ]]; then
    echo -e "${RED}[ERROR]${NC} This script must be run as root (use sudo)" >&2
    exit 1
  fi
}

confirm() {
  read -rp "Are you sure you want to uninstall terrorserver? [y/N] " answer
  case "$answer" in
    [yY][eE][sS]|[yY]) ;;
    *)
      echo "Aborted."
      exit 0
      ;;
  esac
}

stop_service() {
  if systemctl is-active --quiet terror 2>/dev/null; then
    log_info "Stopping terror service..."
    systemctl stop terror
  fi
  if systemctl is-enabled --quiet terror 2>/dev/null; then
    log_info "Disabling terror service..."
    systemctl disable terror
  fi
  if [[ -f "$SERVICE_FILE" ]]; then
    log_info "Removing service file: $SERVICE_FILE"
    rm -f "$SERVICE_FILE"
    systemctl daemon-reload
  fi
}

remove_binary() {
  if [[ -f "$INSTALL_PATH" ]]; then
    log_info "Removing binary: $INSTALL_PATH"
    rm -f "$INSTALL_PATH"
  else
    log_warn "Binary not found at $INSTALL_PATH — skipping"
  fi
}

remove_config() {
  if [[ -d "$CONFIG_DIR" ]]; then
    read -rp "Remove config directory $CONFIG_DIR? [y/N] " answer
    case "$answer" in
      [yY][eE][sS]|[yY])
        log_info "Removing config: $CONFIG_DIR"
        rm -rf "$CONFIG_DIR"
        ;;
      *)
        log_warn "Keeping config at $CONFIG_DIR"
        ;;
    esac
  fi
}

print_done() {
  echo ""
  echo -e "${GREEN}terrorserver has been uninstalled.${NC}"
  echo ""
}

main() {
  require_root
  confirm
  stop_service
  remove_binary
  remove_config
  print_done
}

main "$@"
