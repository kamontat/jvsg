#!/usr/bin/env bash

# Usage:
#   gcp/connect.sh bashclient|goclient|javaclient|server
#   gcp/connect.sh bash|go|java
#   gcp/connect.sh b|g|j|s

TMP="$PWD"
cd "$(dirname "$0")/.." || exit 1
ROOT="$PWD"

# shellcheck disable=SC1091
source "$ROOT/common/base-settings.sh" || exit 2
# shellcheck disable=SC1091
source "$ROOT/common/base-utilities.sh" || exit 2
# shellcheck disable=SC1091
source "$ROOT/common/base-docker.sh" || exit 2

docker_build "$1" "local"

cd "$TMP" || exit 1
unset TMP ROOT
