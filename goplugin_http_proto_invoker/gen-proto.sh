#!/bin/bash

pwd
cd ./models/protos
pwd
protoc --micro_out=../ --go_out=../ prods.proto
cd .. && cd ..
