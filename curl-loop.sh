#!/bin/bash

set -e

if [[ ${1} = "" ]]
then
  echo "Usage: $0 <URL>"
  exit
fi

readonly url=${1}

for i in {0..300}; do
    name=`sed -ne ${RANDOM}p /usr/share/dict/words`
    curl ${url}?name=${name}
done