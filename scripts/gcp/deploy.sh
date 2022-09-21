#!/usr/bin/env bash

# Usage:
#   gcp/deploy.sh bashclient|goclient|javaclient
#   gcp/deploy.sh bash|go|java
#   gcp/deploy.sh b|g|j

TMP="$PWD"
cd "$(dirname "$0")/.." || exit 1
ROOT="$PWD"

# shellcheck disable=SC1091
source "$ROOT/common/base-settings.sh" || exit 2
# shellcheck disable=SC1091
source "$ROOT/common/base-utilities.sh" || exit 2
# shellcheck disable=SC1091
source "$ROOT/common/base-gcloud.sh" || exit 2

gcp_create_instance "$1"

cd "$TMP" || exit 1
unset TMP ROOT
