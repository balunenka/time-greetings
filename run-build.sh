#!/bin/bash

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
PACKAGE_NAME=$(basename "$SCRIPT_DIR")

rm ${PACKAGE_NAME}-*

GOOS=darwin
go build
mv ${PACKAGE_NAME} ${PACKAGE_NAME}-${GOOS}

GOOS=linux
go build
mv ${PACKAGE_NAME} ${PACKAGE_NAME}-${GOOS}

GOOS=windows
go build
mv ${PACKAGE_NAME}.exe ${PACKAGE_NAME}-${GOOS}.exe

zip ${PACKAGE_NAME}.zip ${PACKAGE_NAME}-*