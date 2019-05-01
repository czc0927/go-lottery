#!/bin/bash
kill -9 $(pgrep web)
cd /home/work/server/go/src/lottery
git pull https://github.com/czc0927/go-lottery.git
cd web/
./web &