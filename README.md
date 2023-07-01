

<img src="./assets/mascot.svg" width="150px" align="right">


[![Build Status](https://travis-ci.org/hibernate/hibernate.svg?branch=master)](https://travis-ci.org/hibernate/hibernate)
[![Go Report Card](https://goreportcard.com/badge/github.com/hibernate/hibernate)](https://goreportcard.com/report/github.com/hibernate/hibernate)
[![codecov](https://codecov.io/gh/hibernate/hibernate/branch/master/graph/badge.svg)](https://codecov.io/gh/hibernate/hibernate)
[![GoDoc](https://godoc.org/github.com/hibernate/hibernate?status.svg)](https://godoc.org/github.com/hibernate/hibernate)
![GitHub release](https://img.shields.io/github/release/hibernate/hibernate.svg)


# hibernate - simply save kubernetes resource costs
This is a simple tool to save kubernetes resource costs. It will scale down deployments and statefulsets to 0 replicas at a scheduled time every day. It will also scale up the deployments and statefulsets to the original number of replicas at a scheduled time every day/days of a week.
Primarily built to seemlessly tackle spinning down dev environments at night and weekends.

