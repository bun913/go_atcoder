#!/bin/bash

shopt -s expand_aliases

alias addgo="cp /go/src/work/template.go ./main.go"

contest=$1
if [ $contest = ""]; then
    echo "コンテストidを指定してください"
fi

acc new $contest &&
cd ./$contest/a &&
addgo &&
cd ../b && 
addgo &&
cd ../c &&
addgo &&
cd ../d &&
addgo &&
cd ../e &&
addgo &&
cd ../f &&
addgo
