#!/usr/bin/env bash
chart_dir="./k8s/helm"

if [ -z "$(command -v helm)" ]; then
  echo "$(tput setaf 1)Helm is not installed. This script is for installing the helm chart, not helm itself.$(tput sgr0)"
  exit 1
fi

if [ ! -d "$chart_dir" ]; then
  echo "$(tput setaf 1)Path $chart_dir does not exist, make sure to run the script from project root.$(tput sgr0)"
fi

echo "$(tput setaf 1)$ helm lint $chart_dir$(tput sgr0)"
helm lint "$chart_dir"
echo

if helm list --filter 'aoc2024' | grep -q 'aoc2024'; then
  echo "$(tput setaf 1)Project is already installed, upgrading...$(tput sgr0)"
  helm upgrade aoc2024 "$chart_dir"
else
  echo "$(tput setaf 1)Project is not installed yet, installing...$(tput sgr0)"
  helm install aoc2024 "$chart_dir"
fi