#!/bin/bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
git add .
git commit -m 'test'
git push
