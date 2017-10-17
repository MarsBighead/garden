#!/bin/bash

usage()
{
    echo "usage: $0 -a \"application\" -g \"golang version number\"  -v"
    echo "    -g: golang version,such as 1.8.3"
    echo "    -a: Go application name for build, such as garden"
    echo "    -v: build script version"
 
    exit 1
}
VERSION="0.0.1"
PARSE=`/usr/bin/getopt -q  'a:g:v' "$@"`

if [ $? != 0 ] ; then
    usage
fi

APPLICATION=garden
GOLANG=latest
while [ -n "$1" ] ; do
    case "$1" in
        -a) APPLICATION=$2; shift 2;;
        -g) GOLANG=$2; shift 2 ;;
        -v) echo "$0 $VERSION"; exit 0;;
        --) shift; break ;;
        *) echo "Parameter error"; usage ;;
    esac
done
echo "$APPLICATION $GOLANG"
GOBIN=$GOPATH/bin
GOSRC=$GOPATH/src
APP="/go/src/$APPLICATION/app.sh"
if cid=$(docker ps -a|grep -o -E "build$"); then
    echo "Docker container $cid is removing...."
    docker rm -f $cid
fi
BUILD_CID=$(docker create  -it \
	-v $GOBIN:/go/bin \
	-v $GOSRC:/go/src \
	--name build golang:$GOLANG  \
	/bin/bash $APP )
docker start $BUILD_CID