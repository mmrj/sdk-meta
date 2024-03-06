#!/bin/bash

# This script uses the 'gh' command from Github along with the GraphQL API to enumerate
# all releases for a given repo in the launchdarkly org. The releases are formatted as a JSON array
# suitable for use with jq. Releases < 1.0 are filtered out.
#
# The first arg is the repo name, and the second is an optional prefix to
# strip from the enumerated releases.
#
# You'd only need to pass the second arg if the release tag doesn't exactly match with the
# SDK ID. For example, in the cpp-sdks repo the release tags are prefixed with 'launchdarkly-'.
# To handle this, you'd do:
# $ ./releases.sh cpp-sdks launchdarkly-


if [ -z "$1" ]; then
  echo "Usage: $0 <repo> <optional-prefix-to-remove>"
  exit 1
fi

# TODO: this only fetches up to 100 releases.

# shellcheck disable=SC2016
tags=$(gh api graphql -f query='
query($repo: String!) {
  repository(name: $repo, owner: "launchdarkly") {
    releases(last: 100) {
      nodes {
        tagName
      }
    }
  }
}' -F repo="$1" | jq -c '[.data.repository.releases.nodes[].tagName]' | jq 'map(select(contains("v0") | not))')

if [ -n "$2" ]; then
  echo "${tags//"$2"/}"
else
  echo "$tags"
fi
