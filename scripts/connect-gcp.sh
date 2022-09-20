#!/usr/bin/env bash

gcloud compute \
  ssh \
  --zone "asia-southeast1-a" \
  --project "perfect-crawler-363108" \
  "server"
