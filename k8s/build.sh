#!/usr/bin/env bash
project="$1"
dockerfile="./Dockerfile"

if [ ! -f "$dockerfile" ]; then
  echo "Path $dockerfile does not exist, make sure to run the script from project root."
fi

if [ -z "$(command -v minikube)" ]; then
  echo 'Minikube is not installed, check: https://minikube.sigs.k8s.io/docs/'
  echo 'Build will continue to execute, but you will not be able to use the resulting images in the k8s cluster.'
else
  echo 'Minikube is installed'
  eval $(minikube docker-env)
fi

if [ "$project" = 'all' ]; then
  for day in {01..25}; do
    docker build -t "aoc2024-day$day" --build-arg DAY="$day" --file "$dockerfile" .
  done
else
  docker build -t "aoc2024-$project" --file "$dockerfile" .
fi