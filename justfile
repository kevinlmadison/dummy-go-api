build:
  go build \
    -o server \
    ./cmd/web/

docker:
  docker build . \
    -t "andrewzah/go-test-api"
