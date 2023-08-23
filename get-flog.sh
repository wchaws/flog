#!/usr/bin/env bash
# https://cdn.jsdelivr.net/gh/wchaws/flog/get-flog.sh
set -e

GITHUB_HOST=${GITHUB_HOST:-github.com}
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
URL="https://${GITHUB_HOST}/wchaws/flog/releases/latest/download/flog_${OS}_${ARCH}.tar.gz"

echo "Downloading flog from ${URL}"
curl -SL ${URL} -o /tmp/flog.tar.gz
cd /tmp
tar -xvzf flog.tar.gz
mv -v flog /usr/local/bin
mv -v flog /usr/bin