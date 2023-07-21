#!/bin/bash
export GO111MODULE=on
export GOPROXY=https://goproxy.cn,https://goproxy.io,https://mirrors.aliyun.com/goproxy/,direct
go fmt
make geth 
rm /usr/local/go-ethereum/bin/geth -rf
mv build/bin/geth /usr/local/go-ethereum/bin/
