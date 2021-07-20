#!/bin/bash

cd ../tools

go run tools

cd ..

find resources -name "*.go" -exec goimports -w {} \;
