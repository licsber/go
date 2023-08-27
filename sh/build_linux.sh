#!/usr/bin/env bash

export GOOS=linux

ldflags="
-X 'main.licsber=licsber'
"
cd ../cmd || exit 0
for dir in *; do
  if [ ! -d "$dir" ]; then
    continue
  fi

  name=$(basename "$dir")
  echo "$name"
  cd "$name" || exit 0
  go build -ldflags "$ldflags" -o ../../sh/linux .
  cd .. || exit 0
done

cd ../sh || exit 0
chmod +x linux/*
