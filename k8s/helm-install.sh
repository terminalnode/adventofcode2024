#!/usr/bin/env bash
chart_dir="./k8s/helm"

if [ -z "$(command -v helm)" ]; then
  echo 'Helm is not installed. This script is for installing the helm chart, not helm itself.'
  exit 1
fi

if [ ! -d "$chart_dir" ]; then
  echo "Path $chart_dir does not exist, make sure to run the script from project root."
fi

echo "$ helm lint $chart_dir"
helm lint "$chart_dir"
echo
echo "$ helm install aoc2024 $chart_dir"
helm install aoc2024 "$chart_dir"