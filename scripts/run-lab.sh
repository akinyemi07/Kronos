#!/usr/bin/env bash

set -eu

script_dir="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)"
project_root="$(dirname "$script_dir")"

usage() {
    echo "Usage: $0 <lab-id> [lab-arguments...]" >&2
    echo "Example: $0 linux-timezone Africa/Lagos" >&2
    exit 2
}

if [ "$#" -lt 1 ]; then
    usage
fi

lab_id="$1"
shift

case "$lab_id" in
    linux-timezone)
        lab_script="labs/linux-timezone/check-timezone.sh"
        ;;
    *)
        echo "ERROR: Unknown lab: $lab_id" >&2
        usage
        ;;
esac

lab_script="$project_root/$lab_script"

if [ ! -f "$lab_script" ]; then
    echo "ERROR: Lab script not found: $lab_script" >&2
    exit 1
fi

script_path="$(realpath "$lab_script")"

exec podman run --rm \
    --network=none \
    --volume "$script_path:/lab.sh:ro,Z" \
    docker.io/library/fedora:43 \
    bash /lab.sh "$@"
