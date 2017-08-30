#!/bin/bash

cd /tmp

GOOS=linux GOARCH=amd64 revel package github.com/douban-girls/server

echo "please to /tmp directory to scp this file to your server"