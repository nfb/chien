#!/usr/bin/env bash
cd app
export BINDADDR=:5001
export LOGLEVEL=DEBUG
go run .
