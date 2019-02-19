#!/bin/bash
set -e
basepath=$(cd `dirname $0`; pwd)
cd ${basepath}
protopath=../proto/
cppoutpath=../proto/generatedcpp
if [ ! -d ${cppoutpath} ];then
    mkdir ${cppoutpath}
fi
GRPC_CPP_PLUGIN=grpc_cpp_plugin
GRPC_CPP_PLUGIN_PATH=`which ${GRPC_CPP_PLUGIN}`

protoc -I=$protopath --cpp_out=${cppoutpath} $protopath/ps/*.proto
protoc -I=$protopath --cpp_out=${cppoutpath} $protopath/feeder/*.proto