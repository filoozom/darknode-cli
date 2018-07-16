#!/bin/bash

# Create directory for build files
mkdir build

# Compress darknode directory
zip -r build/darknode.zip darknode

# Copy install script to build folder
cp scripts/install.sh build/install.sh

# Generate binaries
docker-machine create default
docker-machine start
go get github.com/karalabe/xgo
xgo --targets=darwin/amd64,linux/amd64 .
mv darknode-cli-darwin-10.6-amd64 build/darknode_darwin_amd64
mv darknode-cli-linux-amd64 build/darknode_linux_amd64
docker-machine rm -f default