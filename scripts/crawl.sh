#!/bin/bash

set -e

go build ./cmd/ingest

if [ -z "$1" ]; then
  echo "Usage: $0 <sqlite-db-path> <metadata-dir-path>"
  exit 1
fi

golden_db=$1
temp_db="new_$1"
rm -f "$temp_db"

golden_dir=$2
temp_dir="new_$2"
rm -rf "$temp_dir"

sqlite3 "$temp_db" < ./schemas/sdk_metadata.sql
mkdir "$temp_dir"

jq -r '.repos[]' < config.json | while read -r repo; do
  echo "Fetching metadata.json for $repo"
  sanitized_repo=$(echo "$repo" | tr '/' '_')
  gh api "repos/$repo/contents/.launchdarkly/metadata.json" -q '.content' | base64 --decode > "$temp_dir/$sanitized_repo.json"
  echo "Ingesting metadata.json for $repo"
  ./ingest -metadata "$temp_dir/$sanitized_repo.json" -db "$temp_db" -repo "$repo"
done


mv "$temp_db" "$golden_db"

rm -rf "$golden_dir"
mv "$temp_dir" "$golden_dir"
