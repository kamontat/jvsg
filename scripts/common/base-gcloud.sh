#!/usr/bin/env bash

gcp_run() {
  exec_cmd "gcloud" "$@"
}

gcp_connect() {
  local name
  name="$(toapp "$1")"

  gcp_run compute ssh "$name"
}

gcp_create_firewall() {
  local name="$1" port="$2"

  gcloud compute \
    firewall-rules create "$name" \
    --direction=INGRESS \
    --priority=1000 \
    --network=default \
    --action=ALLOW \
    "--rules=tcp:$port,udp:$port" \
    "--target-tags=$name"
}

gcp_create_instance() {
  local name
  name="$(toapp "$1")"

  local cmd zone network_interface maintenance_policy provisioning_model service_account scopes
  cmd="$(getenv "$name" "CREATE_CMD")"
  zone="$(toopt "$name" "zone")"
  image="$(toopt "$name" "image")"
  boot_disk_size="$(toopt "$name" "boot-disk-size")"
  boot_disk_type="$(toopt "$name" "boot-disk-type")"
  boot_disk_device_name="$(toopt "$name" "boot-disk-device-name")"
  container_image="$(toopt "$name" "container-image")"
  container_restart_policy="$(toopt "$name" "container-restart-policy")"
  container_env="$(toopt "$name" "container-env")"
  create_disk="$(toopt "$name" "create-disk")"
  machine_type="$(toopt "$name" "machine-type")"
  network_interface="$(toopt "$name" "network-interface")"
  maintenance_policy="$(toopt "$name" "maintenance-policy")"
  provisioning_model="$(toopt "$name" "provisioning-model")"
  service_account="$(toopt "$name" "service-account")"
  scopes="$(toopt "$name" "scopes")"
  labels="$(toopt "$name" "labels")"
  tags="$(toopt "$name" "tags")"

  # shellcheck disable=SC2207
  params=($(getenv "$name" "CREATE_PARAMETERS"))
  # shellcheck disable=SC2207
  eparams=($(getenv "$name" "CREATE_EPARAMETERS"))

  shift 1
  "gcp_run" compute instances \
    "$cmd" "$name" \
    "$zone" "$image" "$boot_disk_size" \
    "$boot_disk_type" "$boot_disk_device_name" \
    "$container_image" "$container_restart_policy" \
    "$container_env" "$tags" "$create_disk" \
    "$machine_type" "$network_interface" \
    "$maintenance_policy" "$provisioning_model" \
    "$service_account" "$scopes" "$labels" \
    "${params[@]}" "${eparams[@]}" "$@"
}

gcp_update_instance() {
  local name
  name="$(toapp "$1")"

  if [[ "$name" == "server" ]]; then
    echo "Server cannot update via scripts" >&2
    return 1
  else
    image="$(toopt "$name" "container-image")"
    "gcp_run" compute instances \
      update-container \
      "$name" "$image"
  fi
}
