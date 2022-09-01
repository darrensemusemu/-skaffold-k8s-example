# Skaffold k8s Example

This repo is meant to test [`skaffold`](https://skaffold.dev/docs/)'s local dev, builds & deploys.

> Skaffold handles the workflow for building, pushing, and deploying your application, and provides building blocks for creating CI/CD pipelines

## What is needed to run this example:

1. A kubernetes cluster with ingress-nginx, use [kind](https://kind.sigs.k8s.io/) or [minikube](https://minikube.sigs.k8s.io/docs/start/) for local dev
1. Skaffold, read more in the [docs](https://skaffold.dev/docs/) section
1. For deploys, a Kubernetes cluster for deploys (this example uses gke & prow)


### For continuous local kubernetes dev:

Run `skaffold dev`

Update `./cmd/greet-api/main.go` (wait for changes 
to hot reload)

Run `curl -k -i https://localhost:{PORT}/api/skaffold-k8s-example/greet`

### For build & deploy into cluster:

See [./scripts/skaffold-build-deploy.sh](./scripts/skaffold-build-deploy.sh) 

The example server https://playground.darrensemusemu.dev/api/skaffold-k8s-example/greet