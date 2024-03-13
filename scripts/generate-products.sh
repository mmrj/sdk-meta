#!/bin/bash

set -e

sqlite3 -json metadata.sqlite3 "SELECT * from sdk_languages;" |
jq 'reduce .[] as $item ({}; .[$item.id] += [$item.language])' > products/languages.json

sqlite3 -json metadata.sqlite3 "SELECT * from sdk_names;" |
  jq 'reduce .[] as $item ({}; .[$item.id] += [$item.name])' > products/names.json

sqlite3 -json metadata.sqlite3 "SELECT * from sdk_types;" |
  jq 'reduce .[] as $item ({}; .[$item.id] += [$item.type])' > products/types.json

sqlite3 -json metadata.sqlite3 "SELECT * from sdk_repos;" |
  jq 'reduce .[] as $item ({}; .[$item.id] += {github: $item.github})' > products/repos.json

# Do it for sdk_features. Noting that it should contain 'introduced' and 'deprecated' and 'removed' (only if non null in row.)
sqlite3 -json metadata.sqlite3 "SELECT * from sdk_features;" |
  jq 'reduce .[] as $item ({}; .[$item.id] += {introduced: $item.introduced, deprecated: $item.deprecated, removed: $item.removed})' > products/features.json
