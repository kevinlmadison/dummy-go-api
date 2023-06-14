FROM golang:1.20.5-bullseye as build

WORKDIR /build
COPY go.mod go.sum /build/
COPY ./cmd/web/ /build/cmd/web/

WORKDIR /build
RUN go build -o server ./cmd/web/

FROM golang:1.20.5-bullseye

COPY --from=build /build/server /server

ENTRYPOINT ["/server"]
EXPOSE 9090
