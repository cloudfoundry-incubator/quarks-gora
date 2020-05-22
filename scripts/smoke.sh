#!/bin/bash
# Smoke tests for GORA. Meant to be run as a bosh errand
# Requires: curl

ADDRESS=${ADDRESS:-localhost}
SERVER_KEY=${SERVER_KEY:-}
SERVER_CRT=${SERVER_CRT:-}
PORT=${PORT:-4443}
SSL=${SSL:-false}

if [[ "$SSL" == "true" ]]; then
 HOST="--cacert ${SERVER_CRT} -sS https://${ADDRESS}:${PORT}"
else
 HOST="-sS http://${ADDRESS}:${PORT}"
fi

echo
echo "Running GORA smoke tests"
echo
echo "========================="
set -eo pipefail

if curl $HOST | grep -q "SSL=${SSL}"; then
    echo "Can check SSL is present in env var"
else
    echo "SSL not present in env var"
    exit 1
fi

if curl -d "echo FOO;exit 1" $HOST | grep -q "error: FOO"; then
    echo "Correctly returns message on failure"
else
    echo "didn't failed as expected"
    exit 1
fi

if curl -d "echo FOO;exit 0" $HOST | grep -q "OK"; then
    echo "Correctly returns OK when exits successfully"
else
    echo "didn't succeeded as expected"
    exit 1
fi