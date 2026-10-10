#!/usr/bin/env bash

set -eu

echo "===== KRONOS LAB 001 TESTS ====="

output="$(./scripts/run-lab.sh linux-timezone Africa/Lagos)"

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

if ./scripts/run-lab.sh linux-timezone Invalid/NotARealZone; then
    echo "FAIL: Invalid time zone was accepted"
    exit 1
else
    echo "PASS: Invalid time zone rejected"
fi

echo
echo "===== UNKNOWN LAB TEST ====="

if ./scripts/run-lab.sh unknown-lab; then
    echo "FAIL: Unknown lab was accepted"
    exit 1
else
    echo "PASS: Unknown lab rejected"
fi

echo
echo "ALL TESTS PASSED"
