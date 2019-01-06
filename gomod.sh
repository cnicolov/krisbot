#!/bin/bash
set -eu

touch go.mod

CURRENT_DIR=$(basename $(pwd))

CONTENT=$(
  cat <<-EOD
module github.com/cnicolov/${CURRENT_DIR}

require github.com/aws/aws-lambda-go v1.6.0
EOD
)

echo "$CONTENT" >go.mod
