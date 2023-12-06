#!/bin/bash

basepath=$(cd `dirname $0`; pwd)

protoc --go_out=. *.proto
#protoc --proto_path=$basepath/ --go_out=$basepath/ --go_opt=$basepath/ $basepath/*.proto