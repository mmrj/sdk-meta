#!/bin/bash

set -e

base_path="products/json"

# Notes:
# -S argument to JQ is used to sort the keys of the output objects so we get more deterministic output,
# and it's easier to compare diffs between commits to the repo.

sqlite3 -json metadata.sqlite3 "SELECT * from sdk_languages;" |
  jq -S 'reduce .[] as $item ({}; .[$item.id] += [$item.language])' > "$base_path"/languages.json

sqlite3 -json metadata.sqlite3 "SELECT * from sdk_names;" |
  jq -S 'reduce .[] as $item ({}; .[$item.id] = $item.name)' > "$base_path"/names.json

sqlite3 -json metadata.sqlite3 "SELECT * from sdk_types;" |
  jq -S 'reduce .[] as $item ({}; .[$item.id] = $item.type)' > "$base_path"/types.json

sqlite3 -json metadata.sqlite3 "SELECT * from sdk_repos;" |
  jq -S 'reduce .[] as $item ({}; .[$item.id] += {github: $item.github})' > "$base_path"/repos.json

./scripts/eols.sh metadata.sqlite3  |
  jq -n 'reduce inputs[] as $input ({}; .[$input.id] += [$input | del(.id)])' > "$base_path"/releases.json
