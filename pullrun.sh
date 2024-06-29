#!/usr/bin/env bash
git pull
cd app
export BINDADDR=:443
export LOGLEVEL=DEBUG
go run .
