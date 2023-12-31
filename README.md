

<img src="./assets/mascot.svg" width="250px" align="right">


[![Release](https://github.com/Performl/hibernate/actions/workflows/release.yaml/badge.svg)](https://github.com/Performl/hibernate/actions/workflows/release.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/performl/hibernate)](https://goreportcard.com/report/github.com/performl/hibernate)
[![codecov](https://codecov.io/gh/performl/hibernate/branch/master/graph/badge.svg)](https://codecov.io/gh/performl/hibernate)
[![GoDoc](https://godoc.org/github.com/performl/hibernate?status.svg)](https://godoc.org/github.com/performl/hibernate)


# hibernate - simply save kubernetes resource costs
This is a simple tool to save kubernetes resource costs. It will scale down deployments and statefulsets to 0 replicas at a scheduled time every day. It will also scale up the deployments and statefulsets to the original number of replicas at a scheduled time every day/days of a week.
Primarily built to seemlessly tackle spinning down dev environments at night and weekends.

# How it works?
Hibernate simply uses scheduled cronjobs consisting of a go binary which interacts with the kubernetes api to edit the `replica` count for deployments and statefulsets.

When a resource is scaled down to 0, it will store its previous number of replicas in a `ConfigMap` which subsequent wake calls will be able to read. This is to ensure that the resource is scaled up to the previous number of replicas. 

Furthermore, since this is a cronjob based solution, there is no need to run a long running controller process to watch for changes. This is a light weight solution to scale down dev infrastructure costs.

If you do not want to use scheduled cronjobs but would rather use this to automate mass disabling of resources via a single config file, you may opt in to use the binary instead, as the standalone binary can be used as long as you have a valid kubeconfig file.

## Building From Source
This builds the go binary
```bash
git clone git@github.com:Performl/hibernate.git
cd hibernate
make bin
```

## Installing Using Helm
```bash
TODO
```

## Installing Using Kubernetes Manifests
```bash
# prerequisites
kubectl create namespace hibernate

# default manifests
kubectl apply -f https://raw.githubusercontent.com/Performl/hibernate/master/manifests/main.yaml
```

## Dependencies
```bash
# make sure hibernate namespace exists
kubectl create namespace hibernate
```

## Usage (Standalone Binary)
```bash
# to use the binary to sleep resources manually using your local kubernetes config run:
./hibernate --mode=local --action=sleep

# to run in a kubernetes cluster run:
./hibernate --action=sleep

# to wake resources
./hibernate --mode=local --action=sleep
or
./hibernate --action=wake
```

## Authors
* Jacky [@fafnirZ](https://github.com/fafnirZ) [@jxperf](https://github.com/jxperf)