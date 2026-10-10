#!/usr/bin/env bash

set -eu

timezone="${1:-UTC}"

if [ ! -f "/usr/share/zoneinfo/$timezone" ]; then
    echo "ERROR: Unknown time zone: $timezone" >&2
    exit 1
fi

echo "=================================="
echo "KRONOS LAB 001: TIME ZONES"
echo "=================================="

echo
echo "Selected time zone: $timezone"

echo
echo "1. Current UTC time:"
date -u

echo
echo "2. Time in selected time zone:"
TZ="$timezone" date '+%Y-%m-%d %H:%M:%S %Z %z'

echo
echo "3. Container's default time zone:"
date '+%Z %z'

echo
echo "4. System time-zone configuration:"
if [ -L /etc/localtime ]; then
    readlink /etc/localtime
elif [ -e /etc/localtime ]; then
    echo "/etc/localtime exists but is not a symbolic link"
else
    echo "No /etc/localtime file found"
fi

echo
echo "Lab inspection complete."
