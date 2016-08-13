#!/bin/bash

go get github.com/tools/godep

mkdir -p $GOPATH/build

export PATH=$GOPATH/bin:$PATH

# Download the code
cd $GOPATH/src/github.com/cheyang

git clone https://github.com/cheyang/nv-tools.git

#create local cli
cd $GOPATH/src/github.com/cheyang/nv-tools
godep go build -v -ldflags="-s" -o $GOPATH/build/nv-tool

STATUS=${?}

if [[ ${STATUS} -ne 0 ]]; then
  echo "Failed in building nv-tool"
  exit 1
fi

