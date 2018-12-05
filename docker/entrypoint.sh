#!/usr/bin/env bash

echo -e "执行go-script"

nohup /go/src/go-buried-point/docker/main > monitor.log &

go run /go/src/go-buried-point/web-gin/main.go
