#!/bin/bash

POSITIONAL=()
while [[ $# -gt 0 ]]
do
  key="$1"
  case $key in
  -d|--data)
  DATA="$2"
  shift
  shift
    ;;
  -u|--user)
  USER="$2"
  shift
  shift
    ;;
  -p|--password)
  PASSWORD="$2"
  shift
  shift
    ;;
  -o|--organization)
  ORGANIZATION="$2"
  shift
  shift
    ;;
  -db|--database)
  DB="$2"
  shift
  shift
    ;;
  esac
done

DATA="data"
USER="admin"
PASSWORD="password"
ORGANIZATION="hkubx"
DB="F12020"

set -- "${POSITIONAL[@]}"

echo "Running docker command with variables:"
echo "DATA=${DATA}"
echo "USER=${USER}"
echo "PASSWORD=${PASSWORD}"
echo "ORGANIZATION=${ORGANIZATION}"
echo "DB=${DB}"

docker run --name influxdb -m=4gb --cpus=2 -p 8086:8086 -v influxdb_data:/var/lib/influxdb \
  -e INFLUXDB_ADMIN_USER=${USER} \
  -e INFLUXDB_ADMIN_PASSWORD=${PASSWORD} \
  -e INFLUXDB_DB=${DB} \
  --rm quay.io/influxdb/influxdb:v2.0.3