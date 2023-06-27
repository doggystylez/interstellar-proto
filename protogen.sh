#!/bin/bash

[ -z "$(which protoc)" ] && { echo "install protoc" ; exit 1 ; }

for dir in proto/*; do 
    protoc -I ../cosmos-sdk/proto  -I ../gogoproto -I ../cosmos-proto/proto -I ../ibc-go/proto \
        --go_out=. --go-grpc_out=. --proto_path . ${dir}/*
done
cp -r github.com/doggystylez/interstellar-proto/* ./
rm -rf github.com
go mod tidy