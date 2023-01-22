#!/usr/bin/env bash

go install github.com/licsber/go/cmd/query-phone@latest

CMD=~/go/bin/query-phone
input_file='test.txt'
output_file='test.csv'

$CMD >$output_file
while read -r line; do
  line=${line%$'\r'}
  if [[ -z $line ]]; then
    continue
  fi

  $CMD "$line" >>$output_file
done <$input_file

cat $output_file
