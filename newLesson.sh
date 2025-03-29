#!/bin/bash

name="${1}"
path="$(basename "$(dirname "$PWD")")/$(basename "$PWD")"

mkdir $name
cd $name

go mod init "${path}/${name}/main"

touch main.go main_test.go
nvim main_test.go
