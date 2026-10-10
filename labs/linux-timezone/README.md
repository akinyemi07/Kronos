# Linux Lab 001: Time Zones

## Objective

Inspect time-zone configuration and understand the difference between UTC, local time, and system clock synchronization.

## Learning outcomes

By completing this lab, you should be able to:

- Display the current UTC time.
- Inspect the time in a selected time zone.
- Identify the container's default time zone.
- Locate the system time-zone configuration.
- Explain why changing a time zone does not synchronize the system clock.

## Prerequisites

- Podman installed and working.
- The Kronos repository cloned locally.
- Internet access the first time the Fedora container image is downloaded.

## Run the lab

From the root of the Kronos repository, run:

```bash
./scripts/run-lab.sh linux-timezone
```

The default time zone is UTC.

To inspect another time zone, pass its name as an argument:

```bash
./scripts/run-lab.sh linux-timezone Africa/Lagos
```

Other examples include `Europe/London` and `Asia/Tokyo`.

The central runner starts a disposable Fedora container, disables networking, mounts the lab script read-only, and removes the container after execution.

## Run the tests

From the repository root:

```bash
./labs/linux-timezone/test-timezone.sh
```

The test script checks the selected time zone, expected UTC offset, default container time zone, successful completion, and rejection of an invalid time-zone name.

## Tasks

1. Inspect the current UTC time and selected time zone.
2. Find available time zones under `/usr/share/zoneinfo`.
3. Explain the difference between UTC and local time.
4. Explain what NTP does and how synchronization status can be checked.
5. Explain why changing a time zone does not synchronize a clock.

## Safety and limitations

The lab runs in a disposable Fedora container with networking disabled and the lab script mounted read-only.

This is a read-only inspection exercise. It does not change the host's time zone or system clock.

Changing system time zones and managing NTP services require a suitable environment and are outside the scope of this version.
