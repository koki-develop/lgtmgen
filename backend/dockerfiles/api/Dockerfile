FROM golang:1.21 AS builder
WORKDIR /var/task

ENV GOARCH=amd64
ENV GOOS=linux

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download -x

# Build
COPY . .
RUN go build -ldflags="-s -w" -o /var/task/bin/cli .
ENTRYPOINT ["/var/task/bin/cli"]

# ----------------------------

FROM golang:1.21
WORKDIR /var/task
ENV CGO_CFLAGS_ALLOW="-Xpreprocessor"

COPY --from=builder /var/task/bin/cli /var/task/bin/cli
ENTRYPOINT ["/var/task/bin/cli"]
