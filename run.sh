#!/bin/bash

usage="$0 default|updated"

if [[ $# -ne 1 ]]; then
  echo "ERROR: Invalid usage"
  echo "$usage"
  exit 1
fi

config_type="$1"

cp "prometheus.yaml.$config_type" prometheus.yaml
rm -rf "./output/$config_type"

export OUTPUT_DIR="/var/tmp/output/$config_type"

docker-compose build
docker-compose up --force-recreate
