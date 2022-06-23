#!/usr/bin/env bash

function clean()
{
    rm -rf ./bin/*
}

function build()
{
    TARGET="./bin/main"
    BUILD_CMD="go build -o $TARGET"

    for FILE in $(ls -R ./src/*.go); do
        BUILD_CMD="${BUILD_CMD} ${FILE}"
    done

    $BUILD_CMD
}

function run()
{
    ./bin/main
}

case $1 in
    "build")
        build
    ;;
    "run")
        run
    ;;
    "clean")
        clean
    ;;
    *)
        echo "invalid command"
    ;;
esac
