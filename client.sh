#!/usr/bin/env bash
# Simple script to build and run the client module without having to produce
# a binary file
profile_name='aoc2024'
url='localhost'

if [ -n "$(command -v minikube)" ] && minikube profile list | grep -q "$profile_name"; then
  echo "Minikube is installed and profile '$profile_name' exists, using minikube ip ($(minikube ip))."
  echo "Override this by providing your own -url to this command."
  url="$(minikube ip)"
else
  echo "Minikube is not installed or profile '$profile_name' doesn't exist, using localhost."
fi

# shellcheck disable=SC2068
go run ./client -url "$url" $@
