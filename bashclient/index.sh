#!/usr/bin/env bash

SERVER_HOST="${SERVER_HOST:-localhost}"
SERVER_PORT="${SERVER_PORT:-3333}"
SERVER_PATH="${SERVER_PATH:-/json}"
SERVER_DEBUG="${SERVER_DEBUG:-false}"

PUSH_GATEWAY_URL="${PUSH_GATEWAY_URL:-http://localhost:9091}"
JOB_NAME="${JOB_NAME:-bashclient}"

i=0
while true; do
  url="http://$SERVER_HOST:$SERVER_PORT/${SERVER_PATH}?debug=${SERVER_DEBUG}"

  curl -s -X "POST" -w "@formatter.txt" -o /dev/null "$url" |
    curl -s --data-binary @- "$PUSH_GATEWAY_URL/metrics/job/$JOB_NAME"

  echo "request_total $i" |
    curl -s --data-binary @- "$PUSH_GATEWAY_URL/metrics/job/$JOB_NAME"

  ((i++))
  sleep "${EXECUTION_INTERVAL:-1}"
done
