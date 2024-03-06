#!/bin/bash

# This script uses the 'gh' command from Github along with the GraphQL API to enumerate
# all releases for a given repo in the launchdarkly org.
# Releases < 1.0 and alpha/beta tags are filtered out.


if [ -z "$1" ]; then
  echo "Usage: $0 <repo>"
  exit 1
fi

# TODO: this only fetches up to 100 releases.

# shellcheck disable=SC2016
gh api graphql -f query='
query($repo: String!) {
  repository(name: $repo, owner: "launchdarkly") {
    releases(last: 100) {
      nodes {
        tagName
        publishedAt
      }
    }
  }
}' -F repo="$1" | jq -c '[.data.repository.releases.nodes[] |'\
'{tag: .tagName, date: .publishedAt }] | map(select(.tag | contains("v0") | not))'\
' | map(select(.tag | contains("beta") | not))'\
' | map(select(.tag | contains("alpha") | not))'
