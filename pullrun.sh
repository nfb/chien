#!/usr/bin/env bash
git pull
cd app
export BINDPORT=:443
export LOGLEVEL=DEBUG
go run .
