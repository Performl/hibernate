

<img src="./assets/mascot.svg" width="250px" align="right">


[![Build Status](https://travis-ci.org/performl/hibernate.svg?branch=master)](https://travis-ci.org/performl/hibernate)
[![Go Report Card](https://goreportcard.com/badge/github.com/performl/hibernate)](https://goreportcard.com/report/github.com/performl/hibernate)
[![codecov](https://codecov.io/gh/performl/hibernate/branch/master/graph/badge.svg)](https://codecov.io/gh/performl/hibernate)
[![GoDoc](https://godoc.org/github.com/performl/hibernate?status.svg)](https://godoc.org/github.com/performl/hibernate)
![GitHub release](https://img.shields.io/github/release/performl/hibernate.svg)


# hibernate - simply save kubernetes resource costs
This is a simple tool to save kubernetes resource costs. It will scale down deployments and statefulsets to 0 replicas at a scheduled time every day. It will also scale up the deployments and statefulsets to the original number of replicas at a scheduled time every day/days of a week.
Primarily built to seemlessly tackle spinning down dev environments at night and weekends.


## Authors
* [@fafnirZ](https://github.com/fafnirZ)
* [@jxperf](https://github.com/jxperf)