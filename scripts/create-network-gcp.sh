#!/usr/bin/env bash

create_network() {
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

create_network "grafana" "3000"
create_network "prometheus" "9090"
create_network "pushgateway" "9091"
create_network "server" "3333"
