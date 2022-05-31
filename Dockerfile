# Build the manager binary
FROM golang:1.15-alpine as builder

WORKDIR /workspace

# Copy the go source
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o app

FROM alpine
WORKDIR /
COPY --from=builder /workspace/app .

ENTRYPOINT ["/app"]
