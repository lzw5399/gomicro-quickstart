#!/bin/bash

pwd
cd ./models/protos
pwd
protoc --micro_out=../ --go_out=../ prods.proto
protoc-go-inject-tag -input=../prods.pb.go
cd .. && cd ..
