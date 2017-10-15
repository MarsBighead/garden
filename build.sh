#!/bin/bash


GOBIN=$GOPATH/bin
GOSRC=$GOPATH/src
APP=$GOSRC/garden
docker rm build 
MBUILD_CID=$(docker create  -it \
	-v $GOBIN:/go/bin \
	-v $GOSRC:/go/src \
	--name build golang:1.8.3  \
	/bin/bash /go/src/garden/app.sh )

echo $MBUILD_CID
docker start $MBUILD_CID

#m=/Users/paul.duan/go/src/appcoachs.net/x
#docker run --rm \
#   -v $go:$go \
#   -w "$(pwd)" golang:1.8.1 \
#   cd $m/m \
#   go build -v
