#!/usr/bin/env bash

toupper() {
  printf "%s" "$1" | tr '[:lower:]' '[:upper:]'
}

tokey() {
  printf "%s" "$1" | tr '[:lower:]' '[:upper:]' | tr '-' '_'
}

# $1 = app name
# $2 = env name (must be UPPERCASE)
getenv() {
  local defname="DEFAULT"
  local name="$1" key="$2" envname

  envname="$(toupper "$name")_$key"

  eval "local envvalue=\"\$$envname\""
  # shellcheck disable=SC2154
  if test -z "$envvalue"; then
    envname="${defname}_${key}"
    eval "envvalue=\"\$$envname\""
  fi

  # shellcheck disable=SC2154
  printf "%s" "$envvalue"
}

# $1 = app name
# $2 = option key (e.g. 'machine-type' => --machine-type=$machine_type)
# $3 = override variable name (e.g MT => --machine-type=$mt)
toopt() {
  local name="$1" opt="$2" override="$3" key value
  if test -n "$override"; then
    key="$(tokey "$override")"
  else
    key="$(tokey "$opt")"
  fi

  value="$(getenv "$name" "$key")"
  if test -n "$value"; then
    echo "--$opt=$value"
  fi
}

toapp() {
  local name="$1"

  if [[ "$name" == "bashclient" ]] ||
    [[ "$name" == "bash" ]] ||
    [[ "$name" == "b" ]]; then
    printf "%s" "bashclient"
  elif [[ "$name" == "goclient" ]] ||
    [[ "$name" == "go" ]] ||
    [[ "$name" == "g" ]]; then
    printf "%s" "goclient"
  elif [[ "$name" == "javaclient" ]] ||
    [[ "$name" == "java" ]] ||
    [[ "$name" == "j" ]]; then
    printf "%s" "javaclient"
  elif [[ "$name" == "server" ]] ||
    [[ "$name" == "s" ]]; then
    printf "%s" "server"
  else
    printf "%s" "$name"
  fi
}

exec_cmd() {
  local simplified=()

  for i in "$@"; do
    if test -n "$i"; then
      simplified+=("$i")
    fi
  done

  if test -n "$DRYRUN"; then
    echo "$" "${simplified[@]}"
  else
    "${simplified[@]}"
  fi
}
