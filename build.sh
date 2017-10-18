#!/bin/bash

usage()
{
    filename=$(basename $0)
    echo "Usage:   $filename -a \"application\" -g \"golang version\""
    echo "         $filename -h | --help"
    echo "         $filename -v | --version"
    echo "         -g --golang  : Golang version,such as 1.8.3"
    echo "         -a --app     : Go application name for build, such as garden"
    echo "         -v --version : Build script version"
    echo "         -h --help    : Help usage"
 
    exit 1
}
VERSION="0.0.1"
PARSE=`/usr/bin/getopt -q  'a:g:v' --long app:golang:version "$@"`

if [[ $? != 0 ]] || [[ -z $2 ]]; then
    usage
fi

APPLICATION=garden
GOLANG=latest
while [ -n "$1" ] ; do
    case "$1" in
        -a | --app ) APPLICATION=$2; echo "$2";shift 2;;
        -g | --golang) GOLANG=$2; shift 2 ;;
        -v | --version) echo "Version $VERSION"; exit 0;;
        -h | --help) usage;;
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
