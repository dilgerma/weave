#!/bin/bash
export GOPATH=~/go
mkdir $GOPATH
cd $GOPATH
echo "executing go clean"
go clean -i net
echo "executing go install"
go install -tags netgo std
echo "cleaning existing sources"
mkdir -p src/github.com/weaveworks
cd src/github.com/weaveworks
rm -rf weave
echo "cloning"
git clone -b rpi-latest-release http://github.com/dilgerma/weave
cd weave
make SUDO=
