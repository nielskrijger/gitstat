.PHONY: lint goreleaser-dryrun

lint:
	golangci-lint run

goreleaser-dryrun:
	goreleaser --snapshot --skip-publish --rm-dist