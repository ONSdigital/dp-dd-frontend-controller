#!/bin/bash -eux

export BINPATH=$(pwd)/bin
export GOPATH=$(pwd)/go

pushd $GOPATH/src/github.com/ONSdigital/dp-dd-frontend-controller
  go build -o $BINPATH/dp-dd-frontend-controller
popd
