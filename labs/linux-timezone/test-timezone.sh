#!/usr/bin/env bash

set -eu

output="$(podman run --rm \
  --network=none \
  --volume "$PWD/labs/linux-timezone/check-timezone.sh:/lab.sh:ro,Z" \
  docker.io/library/fedora:43 \
  bash /lab.sh)"

printf '%s\n' "$output"

echo
echo "===== VERIFYING LAB RESULTS ====="

if grep -q "KRONOS LAB 001: TIME ZONES" <<< "$output"; then
    echo "PASS: Lab title found"
else
    echo "FAIL: Lab title missing"
    exit 1
fi

if grep -q "Current UTC time:" <<< "$output"; then
    echo "PASS: UTC time check found"
else
    echo "FAIL: UTC time check missing"
    exit 1
fi

if grep -q "UTC +0000" <<< "$output"; then
    echo "PASS: Container reports UTC"
else
    echo "FAIL: Expected UTC configuration not found"
    exit 1
fi

if grep -q "Lab inspection complete." <<< "$output"; then
    echo "PASS: Lab completed"
else
    echo "FAIL: Completion message missing"
    exit 1
fi

echo
echo "ALL CHECKS PASSED"
