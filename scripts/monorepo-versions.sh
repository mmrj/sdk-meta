#!/bin/bash

# This script uses the 'gh' command from Github along with the GraphQL API to enumerate
# all releases for a given monorepo in the launchdarkly org.
# Releases < 1.0 and alpha/beta tags are filtered out.

if [ -z "$1" ]; then
  echo "Usage: $0 <repo> <sdk-id>"
  exit 1
fi

# TODO: this only fetches up to 100 releases.

# shellcheck disable=SC2016
releases=$(./releases.sh "$1")

echo "$releases" | jq 'map(select(.tag | contains($sdk)) | .tag = (.tag | sub("^" + $sub; ""; "")))' --arg sdk "$2" --arg sub "$2-v"
