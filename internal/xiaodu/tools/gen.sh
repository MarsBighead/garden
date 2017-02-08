#!/bin/sh
###############################
#                             #
# Author: Toger               #
# Email : duanmhy@gmail.com   #
# Date  : 2017.02.09          #
#                             #
###############################

OS=`uname`
Platform=`uname -i`
JQ=/usr/bin/jq

cd `pwd`"/../"
currentDir=`pwd`
echo $currentDir
#chmod 666 xiaodu.pb.go
#ls  -lth $currentDir  

#exit
if [[ $OS =~ "Linux" ]]; then
    echo "Start generate protobuf struct under "$OS"..."
    /usr/bin/protoc --go_out=$currentDir xiaodu.proto
    echo "End generate protobuf."
    regLine="import proto \"code.google.com\/p\/goprotobuf\/proto\""
    targetLine="import proto \"github.com\/golang\/protobuf\/proto\""
    Cmd="sed -i '"s/$regLine/$targetLine/g"' xiaodu.pb.go"
    eval $Cmd
    cd txdu
    go test -v xdu_pb.go  xdu_pb_test.go 
    cat xiaodu.json | jq . > t.json 
    mv  t.json xiaodu.json
    cd  ../tools/
    echo `pwd`
else
    echo "os none "$OS
fi
