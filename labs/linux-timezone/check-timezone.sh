#!/usr/bin/env bash

set -eu

echo "=================================="
echo "KRONOS LAB 001: TIME ZONES"
echo "=================================="

echo
echo "1. Current UTC time:"
date -u

echo
echo "2. Current container time zone:"
date +'%Z %z'

echo
echo "3. Time-zone configuration file:"
if [ -L /etc/localtime ]; then
    readlink /etc/localtime
elif [ -e /etc/localtime ]; then
    echo "/etc/localtime exists but is not a symbolic link"
else
    echo "No /etc/localtime file found"
fi

echo
echo "4. Time-zone environment variable:"
printf 'TZ=%s\n' "${TZ:-not set}"

echo
echo "Lab inspection complete."
