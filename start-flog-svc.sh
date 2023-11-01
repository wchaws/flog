#!/usr/bin/env bash
set -eo pipefail

# Install flog
export GITHUB_HOST=ghproxy.com/https://github.com
curl -fsSL https://ghproxy.com/https://github.com/wchaws/flog/blob/master/get-flog.sh | bash

function systemd_conf() {
    local desc="$1"
    local cmd="$2"

    echo "[Unit]
Description=${desc}
Requires=network.target
After=network.target
[Service]
Type=simple
ExecStart=/bin/bash -c \"${cmd}\"
Type=simple
Restart=always
[Install]
WantedBy=multi-user.target"
}

function mksvc() {
    local name="$1"
    local cmd="$2"

    echo "mksvc ${name}"
    echo "$(systemd_conf "${name}" "${cmd}")" > /etc/systemd/system/${name}.service
    systemctl enable ${name}.service
    systemctl start ${name}.service
}

mksvc flog-json "mkdir -p /var/log/json/ && flog -f json -d 2s -l > /var/log/json/access.log"
mksvc flog-nested-json "mkdir -p /var/log/nested-json/ && flog -f nested-json -d 2s -l > /var/log/nested-json/access.log"
mksvc flog-nginx "mkdir -p /var/log/nginx/ && flog -d 2s -l | sed -u -e 's/$/ \\\"-\\\" \\\"curl\/7.79.1\\\" \\\"-\\\"/' > /var/log/nginx/flog-nginx.log"
mksvc flog-apache 'mkdir -p /var/log/apache/ && flog -f apache_common -d 2s -l > /var/log/apache/access.log'
