# Advent of Code 2024
![gofmt](https://github.com/terminalnode/adventofcode2024/workflows/gofmt/badge.svg)
![go test](https://github.com/terminalnode/adventofcode2024/workflows/go%20test/badge.svg)

This year's advent of code is being solved in Go as a series of microservices
running in a Kubernetes cluster (or `docker-compose` if you prefer). The Kubernetes
solution is deployed using Helm and implements service discovery using services and
ingress resources. The `docker-compose` solution also has a form of service discovery
using `traefik`. Both of these will listen on port 80 and expose the individual services
with the prefix `/dayX`.

## Run the project
The project can be run either with Kubernetes (using `minikube`) or with `docker-compose`.
Whichever option you choose, call `./client.sh -file INPUT_FILE -day DAY -part PART` to
run the solution.

If you don't want to use the client, curl is also an option:
`curl -X POST --data-binary @INPUT_FILE http://localhost/DAY/PART`

### Auto-fetch puzzle input
If you copy the file `./client-env.sh.example` to `./client-env.sh` and fill in your
session token, the client (when running `./client.sh`) can auto-fetch your puzzle
input. You can then run the client as only `./client.sh -day DAY -part PART`.

### Running in Kubernetes
* Install `minikube` and run `./k8s/start-minikube.sh` to create the `aoc2024` cluster.
* Then run `./k8s/build.sh all` to build all the docker images, making them available
within `minikube`.
* Then run `./k8s/helm-install.sh` to install or upgrade the Helm chart.

When changing a solution the image needs to be rebuilt, since `helm` has no logic to
build the images for us. Run `./k8s/build.sh dayX` to rebuild a single day, then
rerun `./k8s/helm-install.sh` to reinstall the Helm chart (which will recreate all
deployments).

### Running with `docker-compose`
Just run `docker-compose up` and it will run all services as well as the `traefik`
gateway. If the watch feature of `docker-compose` is enabled, the services will be
automatically rebuilt and redeployed ever time the `common` module or their own
`solutions/dayX` module is updated.

## Progress
* ⭐ means solved
* 🥸 means solved, but takes over a minute to run

| Day | Solution | Day | Solution |
|-----|----------|-----|----------|
| 01  | ⭐ ⭐      | 14  |          |
| 02  | ⭐ ⭐      | 15  |          |
| 03  | ⭐ ⭐      | 16  |          |
| 04  | ⭐ ⭐      | 17  |          |
| 05  | ⭐ ⭐      | 18  |          |
| 06  | ⭐ ⭐      | 19  |          |
| 07  | ⭐ ⭐      | 20  |          |
| 08  | ⭐ ⭐      | 21  |          |
| 09  |          | 22  |          |
| 10  | ⭐ ⭐      | 23  |          |
| 11  | ⭐ ⭐      | 24  |          |
| 12  | ⭐ ⭐      | 25  |          |
| 13  |          |     |          |
