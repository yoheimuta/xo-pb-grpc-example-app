#!/bin/bash

set -euo pipefail

# cd to a BUILD root. {
cd $(dirname $0)
# }

dir=./v1

filenames=(
    user
)

for filename in ${filenames[@]};
do
    protoc \
    -I ${dir} \
    --go_out=../../../../ \
    ${dir}/${filename}.proto
done
