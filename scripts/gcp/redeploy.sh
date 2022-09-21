#!/usr/bin/env bash

# Usage:
#   gcp/redeploy.sh bashclient|goclient|javaclient
#   gcp/redeploy.sh bash|go|java
#   gcp/redeploy.sh b|g|j

TMP="$PWD"
cd "$(dirname "$0")/.." || exit 1
ROOT="$PWD"

# shellcheck disable=SC1091
source "$ROOT/common/base-settings.sh" || exit 2
# shellcheck disable=SC1091
source "$ROOT/common/base-utilities.sh" || exit 2
# shellcheck disable=SC1091
source "$ROOT/common/base-gcloud.sh" || exit 2

gcp_update_instance "$1"

cd "$TMP" || exit 1
unset TMP ROOT
