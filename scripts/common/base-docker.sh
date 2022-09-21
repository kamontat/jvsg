#!/usr/bin/env bash

docker_run() {
  local name localMode="$2"
  name="$(toapp "$1")"

  local params serverHost image
  # shellcheck disable=SC2207
  params=($(getenv "$name" "DOCKER_RUN_PARAMETERS"))
  serverHost="$(getenv "$name" "SERVER_HOST")"
  image="$(getenv "$name" "CONTAINER_IMAGE")"
  if test -n "$localMode"; then
    image="jvsg/$name:latest"
  fi

  go_to_app "$name"
  exec_cmd docker run \
    "-it" "--rm" "${params[@]}" \
    -e "SERVER_HOST=${serverHost}" \
    -e "PUSH_GATEWAY_URL=http://${serverHost}:9091" \
    "$image"
}

docker_build() {
  local name image localMode="$2"
  name="$(toapp "$1")"
  image="$(getenv "$name" "CONTAINER_IMAGE")"
  if test -n "$localMode"; then
    image="jvsg/$name:latest"
  fi

  go_to_app "$name"
  exec_cmd docker build -t "$image" .
}
