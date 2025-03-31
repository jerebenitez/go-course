#!/bin/bash

name="$1"
create_test=true  # Default: create main_test.go

# Parse options
while [[ "$#" -gt 0 ]]; do
    case "$1" in
        --no-test)
            create_test=false
            shift
            ;;
        *)
            name="$1"  # Assume first argument is the project name
            shift
            ;;
    esac
done

mkdir -p "$name"
cd "$name" || exit

touch main.go

# Only create main_test.go if --no-test is not provided
if [ "$create_test" = true ]; then
    touch main_test.go
    nvim main_test.go
else
    nvim main.go
fi

