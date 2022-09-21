#!/usr/bin/env bash

# export DRYRUN="" # Set to true to run in dryrun mode

export GCP_PROJECT_NAME="perfect-crawler-363108"
export DEFAULT_ZONE="asia-southeast1-b"

export SERVER_INTERNAL_IP="10.148.0.8"
export BASHCLIENT_CONTAINER_IMAGE="ghcr.io/kamontat/jvsg-bashclient:sha-28cb04c"
export GOCLIENT_CONTAINER_IMAGE="ghcr.io/kamontat/jvsg-goclient:sha-5fe211d"
export JAVACLIENT_CONTAINER_IMAGE="ghcr.io/kamontat/jvsg-javaclient:sha-9290581"

export BASHCLIENT_CONTAINER_RESTART_POLICY="on-failure"
export BASHCLIENT_CONTAINER_ENV="SERVER_HOST=$SERVER_INTERNAL_IP,PUSH_GATEWAY_URL=http://$SERVER_INTERNAL_IP:9091"
export BASHCLIENT_IMAGE="projects/cos-cloud/global/images/cos-stable-101-17162-40-1"
export BASHCLIENT_BOOT_DISK_SIZE="10GB"
export BASHCLIENT_BOOT_DISK_TYPE="pd-balanced"
export BASHCLIENT_BOOT_DISK_DEVICE_NAME="bashclient"
export BASHCLIENT_CREATE_EPARAMETERS="--container-privileged"
export BASHCLIENT_LABELS="container-vm=cos-stable-101-17162-40-1"

export GOCLIENT_CONTAINER_RESTART_POLICY="on-failure"
export GOCLIENT_CONTAINER_ENV="SERVER_HOST=$SERVER_INTERNAL_IP,PUSH_GATEWAY_URL=http://$SERVER_INTERNAL_IP:9091"
export GOCLIENT_IMAGE="projects/cos-cloud/global/images/cos-stable-101-17162-40-1"
export GOCLIENT_BOOT_DISK_SIZE="10GB"
export GOCLIENT_BOOT_DISK_TYPE="pd-balanced"
export GOCLIENT_BOOT_DISK_DEVICE_NAME="goclient"
export GOCLIENT_CREATE_EPARAMETERS="--container-privileged"
export GOCLIENT_LABELS="container-vm=cos-stable-101-17162-40-1"

export JAVACLIENT_CONTAINER_RESTART_POLICY="on-failure"
export JAVACLIENT_CONTAINER_ENV="SERVER_HOST=$SERVER_INTERNAL_IP,PUSH_GATEWAY_URL=http://$SERVER_INTERNAL_IP:9091"
export JAVACLIENT_IMAGE="projects/cos-cloud/global/images/cos-stable-101-17162-40-1"
export JAVACLIENT_BOOT_DISK_SIZE="10GB"
export JAVACLIENT_BOOT_DISK_TYPE="pd-balanced"
export JAVACLIENT_BOOT_DISK_DEVICE_NAME="javaclient"
export JAVACLIENT_CREATE_EPARAMETERS="--container-privileged"
export JAVACLIENT_LABELS="container-vm=cos-stable-101-17162-40-1"

export SERVER_CREATE_CMD="create"
export SERVER_MACHINE_TYPE="e2-custom-4-4096"
export SERVER_TAGS="grafana,prometheus,server,pushgateway"
export SERVER_CREATE_DISK="auto-delete=yes,boot=yes,device-name=server,image=projects/debian-cloud/global/images/debian-11-bullseye-v20220822,mode=rw,size=10,type=projects/$GCP_PROJECT_NAME/zones/us-west4-b/diskTypes/pd-balanced"
export SERVER_CREATE_EPARAMETERS="--reservation-affinity=any"

export DEFAULT_SERVER_HOST="$SERVER_INTERNAL_IP"

export DEFAULT_CREATE_CMD="create-with-container"

export DEFAULT_NETWORK_INTERFACE="network-tier=PREMIUM,subnet=default"

export DEFAULT_MAINTENANCE_POLICY="MIGRATE"

export DEFAULT_PROVISIONING_MODEL="STANDARD"

export DEFAULT_SERVICE_ACCOUNT="936767540053-compute@developer.gserviceaccount.com"

export DEFAULT_SCOPES="https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring.write,https://www.googleapis.com/auth/servicecontrol,https://www.googleapis.com/auth/service.management.readonly,https://www.googleapis.com/auth/trace.append"

export DEFAULT_CREATE_PARAMETERS="--no-shielded-secure-boot --shielded-vtpm --shielded-integrity-monitoring"

export DEFAULT_CREATE_EPARAMETERS=""
