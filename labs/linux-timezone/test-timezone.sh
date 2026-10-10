#!/usr/bin/env bash

set -eu

echo "===== RUNNING KRONOS LAB 001 ====="

output="$(podman run --rm \
  --network=none \
  --volume "$PWD/labs/linux-timezone/check-timezone.sh:/lab.sh:ro,Z" \
  docker.io/library/fedora:43 \
  bash /lab.sh)"

printf '%s\n' "$output"

echo
echo "===== VERIFYING RESULTS ====="

check() {
    local description="$1"
    local pattern="$2"

    if grep -Fq "$pattern" <<< "$output"; then
        echo "PASS: $description"
    else
        echo "FAIL: $description"
        exit 1
    fi
}

check "Lab title" "KRONOS LAB 001: TIME ZONES"
check "UTC time inspection" "Current UTC time:"
check "UTC time-zone offset" "UTC +0000"
check "UTC configuration path" "/usr/share/zoneinfo/Etc/UTC"
check "Lab completion" "Lab inspection complete."

echo
echo "ALL CHECKS PASSED"
