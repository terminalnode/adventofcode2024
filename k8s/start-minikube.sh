#!/usr/bin/env bash
# Small script for setting up a cluster using Minikube,
# if a cluster is already set up, it will start/switch to the
# correct profile. The profile is 'aoc2024'.
profile_name='aoc2024'

if [ -z "$(command -v minikube)" ]; then
  echo 'Minikube is not installed, check: https://minikube.sigs.k8s.io/docs/'
  exit 1
fi
if [ -z "$(command -v kubectl)" ]; then
  echo 'Kubectl is not installed, this should come with Minikube I think?'
  exit 1
fi

if minikube profile list | grep -q "$profile_name"; then
  if minikube status -p "$profile_name" | grep -q "Stopped"; then
    echo "Profile exists but is stopped, starting..."
    minikube start -p "$profile_name"
    echo "...and activating."
    minikube profile "$profile_name"
  else
    echo "Profile exists and is started! Activating..."
    minikube profile "$profile_name"
  fi
else
  echo "Profile $profile_name does not exist, creating..."
  minikube start -p "$profile_name"
  echo "...and activating..."
  minikube profile "$profile_name"
  echo "...and activating the ingress controller addon (nginx ingress controller)."
  minikube addons enable ingress
fi

minikube profile list
echo "Your current active kubernetes context: $(kubectl config current-context)"