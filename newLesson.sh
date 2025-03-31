#!/bin/bash

name="${1}"
# path="$(basename "$(dirname "$PWD")")/$(basename "$PWD")"

mkdir $name
cd $name

# No longer needed, since I no longer create multiple modules
# go mod init "${path}/${name}/main"

touch main.go main_test.go
nvim main_test.go
