#!/usr/bin/env bash
set -euo pipefail

# ─────────────────────────────────────────────
# terrorserver uninstall.sh
# ─────────────────────────────────────────────

INSTALL_PATH="/usr/local/bin/terror"
CONFIG_DIR="/etc/terror"
SERVICE_FILE="/etc/systemd/system/terror.service"
SERVICE_NAME="terror"

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

require_systemctl() {
  if ! command -v systemctl &>/dev/null; then
    log_error "systemctl is required to uninstall the terror systemd service"
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
  if systemctl is-active --quiet "$SERVICE_NAME" 2>/dev/null; then
    log_info "Stopping $SERVICE_NAME service..."
    systemctl stop "$SERVICE_NAME"
  else
    log_warn "$SERVICE_NAME service is not running — skipping stop"
  fi

  if systemctl is-enabled --quiet "$SERVICE_NAME" 2>/dev/null; then
    log_info "Disabling $SERVICE_NAME service..."
    systemctl disable "$SERVICE_NAME"
  else
    log_warn "$SERVICE_NAME service is not enabled — skipping disable"
  fi

  if systemctl list-unit-files "$SERVICE_NAME.service" --no-legend 2>/dev/null | grep -q "$SERVICE_NAME.service"; then
    log_info "Resetting failed state for $SERVICE_NAME service..."
    systemctl reset-failed "$SERVICE_NAME" 2>/dev/null || true
  fi

  if [[ -f "$SERVICE_FILE" ]]; then
    log_info "Removing service file: $SERVICE_FILE"
    rm -f "$SERVICE_FILE"
  else
    log_warn "Service file not found at $SERVICE_FILE — skipping"
  fi

  systemctl daemon-reload
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
  else
    log_warn "Config directory not found at $CONFIG_DIR — skipping"
  fi
}

print_done() {
  echo ""
  echo -e "${GREEN}terrorserver has been uninstalled.${NC}"
  echo ""
}

main() {
  require_root
  require_systemctl
  confirm
  stop_service
  remove_binary
  remove_config
  print_done
}

main "$@"
