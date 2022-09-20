#!/usr/bin/env bash

create_instance() {
  local name="$1" image="$2" server="$3"

  gcloud compute instances \
    create-with-container \
    "$name" \
    --zone=asia-southeast1-a \
    --machine-type=e2-micro \
    --network-interface=network-tier=PREMIUM,subnet=default \
    --maintenance-policy=MIGRATE \
    --provisioning-model=STANDARD \
    --service-account=936767540053-compute@developer.gserviceaccount.com \
    --scopes=https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring.write,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/trace.append \
    --image=projects/cos-cloud/global/images/cos-stable-101-17162-40-1 \
    --boot-disk-size=10GB \
    --boot-disk-type=pd-balanced \
    "--boot-disk-device-name=$name" \
    "--container-image=$image" \
    --container-restart-policy=on-failure \
    --container-privileged \
    "--container-env=SERVER_HOST=$server,PUSH_GATEWAY_URL=http://$server:9091" \
    --no-shielded-secure-boot \
    --shielded-vtpm \
    --shielded-integrity-monitoring \
    --labels=container-vm=cos-stable-101-17162-40-1
}

create_instance \
  "goclient" \
  "ghcr.io/kamontat/jvsg-goclient:sha-b9f6d08" \
  "10.148.0.8"
