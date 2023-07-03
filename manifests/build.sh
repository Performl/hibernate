#!/bin/bash

# recreates main.yaml
rm main.yaml
touch main.yaml

# concatenates all manifests inside independent/ into a single file
for i in independent/*.yaml; do
  cat $i >> main.yaml
  echo "\n" >> main.yaml # ensures whitespace
  echo "---" >> main.yaml
done
