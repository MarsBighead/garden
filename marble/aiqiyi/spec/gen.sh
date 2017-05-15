#!/bin/sh
Machine=`uname -i`
OutputDir=`pwd`"/.."
UnionAiqiyi="aiqiyi.proto"
RequestAiqiyi="bid_request.proto"
ResponseAiqiyi="bid_response.proto"
pkgName="aiqiyi.pb.go"
function fixPackage() {
    regxL1="const _ = proto.ProtoPackageIsVersion2 \/\/ please upgrade the proto package"
    regxL2="func init() { proto.RegisterFile(\"aiqiyi.proto\", fileDescriptor0) }"
	if [[ !(-f $1) ]];then
		echo "$1 not exist, exiting..."
		exit
	fi
	if [[  $Machine =~ "Mac" ]];then
        echo "Running on Mac to edit $1 with sed."
        #pwd
        sed -i "" "s/$regxL1/\/\/ $regxL1/g" $1
        sed -i "" "s/$regxL2/\/\/ $regxL2/g" $1
    else 
        echo "Running on machine $Machine to edit $1 with sed."
        sed -i  "s/$regxL1/\/\/ $regxL1/g" $1
  		sed -i  "s/$regxL2/\/\/ $regxL2/g" $1
    fi
}


# Insert the following 2 lines to define
# protobuf version and package information
#     syntax = "proto2";
#     package aiqiyi; 
# and delete lines
#     with extensions 100 to max;
cat $RequestAiqiyi | sed   '/package ads_serving.proto;/i\
syntax = \"proto2\";\
package aiqiyi;\
' | sed "5d" | sed /"extensions 100 to max;"/d> $OutputDir"/"$UnionAiqiyi
if  [ -f $OutputDir"/"$UnionAiqiyi ]; then
    echo "Generating $UnionAiqiyi request part successufflly!"
    cat $ResponseAiqiyi | sed '1,4d'  | sed /"extensions 100 to max;"/d >>  $OutputDir"/"$UnionAiqiyi
    cd $OutputDir
    echo "Start generate package with .proto file $UnionAiqiyi."
    protoc  --plugin=$GOPATH/bin/protoc-gen-go --go_out=$OutputDir $UnionAiqiyi
	if [[ -f $pkgName ]]; then
    	echo "Start fix generated go package file aiqiyi.pb.go."
		fixPackage $pkgName
	fi
else
    echo "Generating $UnionAiqiyi request part failed! Exiting..." 
    exit
fi
