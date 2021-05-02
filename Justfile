_help:
    @just --list --unsorted

test:
    go test -cover ./...

upgrade:
    go get github.com/crdx/assert
    go mod tidy
