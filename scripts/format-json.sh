#!/bin/bash

for file in ./data/*.json; do
  jq . "$file" > "$file.tmp" && mv "$file.tmp" "$file"
done
