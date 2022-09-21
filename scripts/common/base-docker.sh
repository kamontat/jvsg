#!/usr/bin/env bash

docker_run() {
  local name
  name="$(toapp "$1")"

  # shellcheck disable=SC2207
  params=($(getenv "$name" "DOCKER_RUN_PARAMETERS"))
  serverHost="$(getenv "$name" "SERVER_HOST")"
  image="$(getenv "$name" "CONTAINER_IMAGE")"

  exec_cmd docker run \
    "-it" "--rm" "${params[@]}" \
    -e "SERVER_HOST=${serverHost}" \
    -e "PUSH_GATEWAY_URL=http://${serverHost}:9091" \
    "$image"
}
