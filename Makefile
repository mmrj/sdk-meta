.PHONY: crawl products gen-code-schema

# Crawl all the repos and update metadata.sqlite3 with the results.
# GITHUB_TOKEN must be set in the environment.
crawl:
	./scripts/crawl.sh metadata.sqlite3 metadata

# Generate all the JSON products.
products:
	./scripts/generate-products.sh

all: crawl products

gen-code-schema:
	go-jsonschema -p logs logs/schema/ld_log_codes.json > lib/logs/ld_log_codes.go
	# This is to work around an inadequacy in the generated types in how it handles patternProperties.
	sed -i '/type LdLogCodesJsonConditions map\[string\]interface{}/ c\type LdLogCodesJsonConditions map\[string\]Condition' lib/logs/ld_log_codes.go