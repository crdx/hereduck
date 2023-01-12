@_help:
    just --list --unsorted

# run tests
@test *args:
    cd src && go test -cover ./... {{ args }}
