#!/usr/bin/env bash

# Usage:
#   gcp/firewall.sh "<name>" "<port>"  - Create custom rule
#   gcp/firewall.sh                    - Create all predefined rules

TMP="$PWD"
cd "$(dirname "$0")/.." || exit 1
ROOT="$PWD"

# shellcheck disable=SC1091
source "$ROOT/common/base-settings.sh" || exit 2
# shellcheck disable=SC1091
source "$ROOT/common/base-utilities.sh" || exit 2
# shellcheck disable=SC1091
source "$ROOT/common/base-gcloud.sh" || exit 2

if [[ $# -ge 2 ]]; then
  gcp_create_firewall "$1" "$2"
else
  gcp_create_firewall "grafana" "3000"
  gcp_create_firewall "prometheus" "9090"
  gcp_create_firewall "pushgateway" "9091"
  gcp_create_firewall "server" "3333"
fi

cd "$TMP" || exit 1
unset TMP ROOT
