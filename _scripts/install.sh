#!/bin/bash

BASE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
echo "$BASE_DIR"
SERVICE="baseballgame"
SERVICE_DIR="$BASE_DIR/../$SERVICE-service"
UPSTART="$SERVICE.conf"

if [[ ! -f "/etc/init/$UPSTART" ]]; then
  echo "/etc/init/$UPSTART doesn't exist, create it"
  if [[ -f "$BASE_DIR/upstart.conf" ]]; then
    printf "$BASE_DIR/upstart.conf found, creating..."
    cp "$BASE_DIR/upstart.conf" "/etc/init/$UPSTART"
    chmod 0700 "/etc/init/$UPSTART"
    initctl list
    printf "done.\n"
  else
    echo "$BASE_DIR/upstart.conf file not found. unable to create upstart script"
  fi
else
  echo "/etc/init/$UPSTART found. skipping creation."
fi

if [[ ! -d "$SERVICE_DIR" ]]; then
  echo "creating service directory: $SERVICE_DIR"
  mkdir -p "$SERVICE_DIR"
fi

rm -f "$SERVICE_DIR/$SERVICE"
cp "$BASE_DIR/$SERVICE" "$SERVICE_DIR/."

service "$SERVICE" restart
printf "Status: "
initctl list | grep "$SERVICE"

cd "$BASE_DIR/.."
rm -f "\*-$SERVICE_linux\*"
rm -fR "$BASE_DIR"