#!/bin/bash

set -euo pipefail

# cd to a BUILD root. {
cd $(dirname $0)
# }

dir=./v1

filenames=(
    user
    userproduct
)

for filename in ${filenames[@]};
do
    protoc \
    -I ${dir} \
    --go_out=plugins=grpc:../../../../ \
    ${dir}/${filename}.proto
done
