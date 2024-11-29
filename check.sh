#!/usr/bin/env bash
bpftool map show | grep xaos
bpftool prog show | grep rock
ls /sys/fs/bpf
