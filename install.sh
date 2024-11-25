#!/bin/env bash

set -eo pipefail

go build -o /usr/local/bin/xray-builder

ETC_XRAYBUILDER=/usr/local/etc/xray-builder
mkdir -p $ETC_XRAYBUILDER
cp -r ./configs/* $ETC_XRAYBUILDER
