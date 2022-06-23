#!/usr/bin/env bash

function build()
{
    TARGET="./bin/main"
    BUILD_CMD="go build -o $TARGET"

    for FILE in $(ls -R ./src/*.go); do
        BUILD_CMD="${BUILD_CMD} ${FILE}"
    done

    $BUILD_CMD
}

case $1 in
    "build")
        build
    ;;
    "run")
        ./bin/main
    ;;
    "clean")
        rm -rf ./bin/*
    ;;
    *)
        echo "invalid command"
    ;;
esac
