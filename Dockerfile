FROM ossys/golang:1.19.7-fips-rocky

WORKDIR /build
COPY go.mod go.sum /build/
COPY ./cmd/web/ /build/cmd/web/

WORKDIR /build
RUN go build -o server ./cmd/web/

ENTRYPOINT ["/build/server"]
EXPOSE 9090
