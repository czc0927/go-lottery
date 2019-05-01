#!/bin/bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
cd ../
git config  user.name "czc0927"
git config  user.email "963239044@qq.com"
git add .
git commit -m 'test'
git push
