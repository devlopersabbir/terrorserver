#!/usr/bin/env bash
set -euo pipefail

# ─────────────────────────────────────────────
# terrorserver uninstall.sh
# ─────────────────────────────────────────────

INSTALL_PATH="/usr/local/bin/terror"
CONFIG_DIR="/etc/terror"
CONFIG_FILE="$CONFIG_DIR/Runtime"
WEB_ROOT="${TERROR_WEB_ROOT:-/var/www/terrorserver}"
CERT_CACHE="${TERROR_CERT_CACHE:-/var/lib/terror/certs}"
SERVICE_FILE="/etc/systemd/system/terror.service"
SERVICE_NAME="terror"
AUTHOR_NAME="Sabbir Hossain Shuvo"
AUTHOR_URL="https://devlopersabbir.github.io"
FORCE_YES=false


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
  if [[ "$FORCE_YES" == "true" ]]; then
    return
  fi

  local CYAN='\033[0;36m'
  local GRAY='\033[0;90m'
  local WHITE='\033[1;37m'
  local BOLD='\033[1m'

  echo -e ""
  echo -e "  ${RED}${BOLD}▲ TERRORSERVER${NC} ${GRAY}uninstall${NC}"
  echo -e "  ${WHITE}${BOLD}TARGETS${NC}"
  echo -e "  ${GRAY}├─${NC} ${BOLD}Binary:${NC}   ${GRAY}${INSTALL_PATH}${NC}"
  echo -e "  ${GRAY}├─${NC} ${BOLD}Config:${NC}   ${GRAY}${CONFIG_FILE}${NC}"
  echo -e "  ${GRAY}├─${NC} ${BOLD}Web Root:${NC} ${GRAY}${WEB_ROOT}${NC}"
  echo -e "  ${GRAY}└─${NC} ${BOLD}Certs:${NC}    ${GRAY}${CERT_CACHE}${NC}"
  echo -e ""
  
  # Read from /dev/tty to avoid consuming the script when piped from curl
  read -rp "$(echo -e "  ${CYAN}Continue uninstall?${NC} [y/N] ")" answer < /dev/tty || answer="n"
  
  case "$answer" in
    [yY][eE][sS]|[yY]) ;;
    *)
      echo -e ""
      echo -e "  ${YELLOW}[WARN]${NC} Uninstall aborted."
      echo -e ""
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

  # Also kill any lingering processes not managed by systemd
  if pgrep -x "$SERVICE_NAME" >/dev/null; then
    log_info "Killing lingering $SERVICE_NAME processes..."
    pkill -9 "$SERVICE_NAME" || true
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
    local answer
    if [[ "$FORCE_YES" == "true" ]]; then
      answer="y"
    else
      read -rp "Remove config directory $CONFIG_DIR? [y/N] " answer < /dev/tty || answer="n"
    fi

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

remove_welcome_site() {
  if [[ -d "$WEB_ROOT" ]]; then
    local answer
    if [[ "$FORCE_YES" == "true" ]]; then
      answer="y"
    else
      read -rp "Remove welcome site directory $WEB_ROOT? [y/N] " answer < /dev/tty || answer="n"
    fi

    case "$answer" in
      [yY][eE][sS]|[yY])
        log_info "Removing welcome site: $WEB_ROOT"
        rm -rf "$WEB_ROOT"
        ;;
      *)
        log_warn "Keeping welcome site at $WEB_ROOT"
        ;;
    esac
  else
    log_warn "Welcome site directory not found at $WEB_ROOT — skipping"
  fi
}

remove_cert_cache() {
  local cert_root
  cert_root="$(dirname "$CERT_CACHE")"

  if [[ -d "$cert_root" ]]; then
    local answer
    if [[ "$FORCE_YES" == "true" ]]; then
      answer="y"
    else
      read -rp "Remove certificate cache directory $cert_root? [y/N] " answer < /dev/tty || answer="n"
    fi

    case "$answer" in
      [yY][eE][sS]|[yY])
        log_info "Removing certificate cache: $cert_root"
        rm -rf "$cert_root"
        ;;
      *)
        log_warn "Keeping certificate cache at $cert_root"
        ;;
    esac
  else
    log_warn "Certificate cache not found at $cert_root — skipping"
  fi
}

print_done() {
  local CYAN='\033[0;36m'
  local GRAY='\033[0;90m'
  local WHITE='\033[1;37m'
  local BOLD='\033[1m'

  echo -e ""
  echo -e "  ${RED}${BOLD}▲ TERRORSERVER${NC} ${GRAY}removed${NC}"
  echo -e "  ${GREEN}Successfully uninstalled from your system${NC}"
  echo -e ""
  echo -e "  ${WHITE}${BOLD}SUMMARY${NC}"
  echo -e "  ${GRAY}├─${NC} ${BOLD}Service:${NC}  ${CYAN}${SERVICE_NAME}${NC} stopped and disabled"
  echo -e "  ${GRAY}├─${NC} ${BOLD}Binary:${NC}   ${GRAY}${INSTALL_PATH}${NC}"
  echo -e "  ${GRAY}├─${NC} ${BOLD}Config:${NC}   ${GRAY}${CONFIG_DIR}${NC}"
  echo -e "  ${GRAY}├─${NC} ${BOLD}Web Root:${NC} ${GRAY}${WEB_ROOT}${NC}"
  echo -e "  ${GRAY}└─${NC} ${BOLD}Certs:${NC}    ${GRAY}${CERT_CACHE}${NC}"
  echo -e ""
  echo -e "  ${WHITE}${BOLD}MAINTAINER${NC}"
  echo -e "  ${GRAY}By${NC} ${AUTHOR_NAME} ${GRAY}(${AUTHOR_URL})${NC}"
  echo -e "  ${GRAY}Docs: https://terror.softvenceomega.com/docs${NC}"
  echo -e ""
}

main() {
  for arg in "$@"; do
    case "$arg" in
      -y|--yes) FORCE_YES=true ;;
    esac
  done

  require_root
  require_systemctl
  confirm
  stop_service
  remove_binary
  remove_config
  remove_welcome_site
  remove_cert_cache
  print_done
}

main "$@"
