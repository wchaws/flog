#!/usr/bin/env bash
set -e

GITHUB_HOST=${GITHUB_HOST:-github.com}
FLOG_VERSION=${FLOG_VERSION}
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)
case $ARCH in
    armv5*) ARCH="armv5";;
    armv6*) ARCH="armv6";;
    armv7*) ARCH="armv7";;
    aarch64) ARCH="arm64";;
    x86) ARCH="386";;
    x86_64) ARCH="amd64";;
    i686) ARCH="386";;
    i386) ARCH="386";;
esac

curl -LS "https://${GITHUB_HOST}/wchaws/flog/releases/download/v${FLOG_VERSION}/flog_${FLOG_VERSION}_${OS}_${ARCH}.tar.gz" -o /tmp/flog.tar.gz
cd /tmp
tar -xvzf flog.tar.gz
mv flog /usr/bin