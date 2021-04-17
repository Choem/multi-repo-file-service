#!/usr/bin/env bash

# Get project root path
ROOT_PATH=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )/..

# Load libs
source ${ROOT_PATH}/scripts/load_env.sh

# Minio secret
kubectl create secret generic minio-secret \
  --from-literal=accesskey=$MINIO_ACCESS_KEY \
  --from-literal=secretkey=$MINIO_SECRET_KEY