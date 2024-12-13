#!/usr/bin/env bash
# Creates a certificate for the Kubernetes ingress controller,
# this is required by the nginx-ingress controller to enable HTTP/2,
# and HTTP 2 is required for gRPC.
#
# The certificate is already created and checked into this repository,
# but the command is left here as documentation. The checked in certificate
# is valid for ~100 years.

openssl req -x509 -nodes -days 36500 -newkey rsa:2048 \
  -keyout k8s/helm/certs/aoc2024.key -out k8s/helm/certs/aoc2024.crt \
  -subj "/CN=*.aoc2024.se/O=AdventOfCode" \
  -config ./k8s/cert.config
