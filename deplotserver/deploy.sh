#!/bin/bash
su -
kill -9 $(pgrep web)
cd /home/work/server/go/src/lottery
git clean  -d  -fx
git pull https://github.com/czc0927/go-lottery.git
cd /home/work/server/go/src/lottery/web/
nohup ./web >> ./output.log 2>&1 &
#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build