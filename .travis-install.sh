#! /usr/bin/env sh
echo install
mkdir -p output/strict
go env
go get -t -v ./...
#export RMANTREE=./opt
#export RIGO2_DEBUG="testing"
