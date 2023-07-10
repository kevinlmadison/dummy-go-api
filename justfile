build:
  go build \
    -o server \
    ./cmd/web/

docker:
  docker build . \
    -t "andrewzah/go-test-api:0.0.1"

ci:
  act --secret-file ci-secrets.env
