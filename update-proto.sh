#!/bin/bash

# TODO: versioning
rm -rf api/helium/*
curl -s -L https://github.com/helium/proto/archive/refs/heads/master.tar.gz | tar --strip-components=2 -C ./api/helium -xzf - proto-master/src

echo "----------------------------------------------------------------------------------------------------------------"
echo "Please carefully review the changes and re-apply any fixes that we need in order to make this work for golang!"
echo "After that you can run build.sh to generate everything from the protobuf definitions + build the commands"
echo "----------------------------------------------------------------------------------------------------------------"
