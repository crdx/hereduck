set quiet := true

[private]
help:
    just --list --unsorted

test:
    go test -cover ./...

fmt:
    just --fmt
