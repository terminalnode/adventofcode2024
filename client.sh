#!/usr/bin/env bash
# Simple script to build and run the client module without having to produce
# a binary file

# shellcheck disable=SC2068
go run ./client $@
