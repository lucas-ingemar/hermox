# development stage
FROM golang:1.25 AS dev
WORKDIR /src

# download modules
COPY go.* ./
COPY cmd/ ./cmd
RUN go mod download
RUN go build -o hermox ./cmd/hermox/main.go

# release stage
FROM alpine AS release
COPY --from=dev /src/hermox /usr/local/bin/hermox
ENTRYPOINT ["/usr/local/bin/hermox"]
