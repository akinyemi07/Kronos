#!/usr/bin/env bash

set -eu

echo "===== KRONOS LAB 001 TESTS ====="

output="$(podman run --rm \
  --network=none \
  --volume "$PWD/labs/linux-timezone/check-timezone.sh:/lab.sh:ro,Z" \
  docker.io/library/fedora:43 \
  bash /lab.sh Africa/Lagos)"

printf '%s\n' "$output"

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

echo
echo "===== VALID TIME-ZONE TESTS ====="

check "Selected time zone is Lagos" \
    "Selected time zone: Africa/Lagos"

check "Lagos offset is +0100" \
    "WAT +0100"

check "Container default remains UTC" \
    "UTC +0000"

check "Lab completes successfully" \
    "Lab inspection complete."

echo
echo "===== INVALID TIME-ZONE TEST ====="

if podman run --rm \
    --network=none \
    --volume "$PWD/labs/linux-timezone/check-timezone.sh:/lab.sh:ro,Z" \
    docker.io/library/fedora:43 \
    bash /lab.sh Invalid/NotARealZone
then
    echo "FAIL: Invalid time zone was accepted"
    exit 1
else
    echo "PASS: Invalid time zone rejected"
fi

echo
echo "ALL TESTS PASSED"
