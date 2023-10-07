ARG BASE_IMAGE

# ----------------------------
FROM ${BASE_IMAGE} AS builder

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download -x

# Build
COPY . .
RUN go build -ldflags="-s -w" -o /var/task/bin/cli .

# ----------------------------

FROM ${BASE_IMAGE}

COPY --from=builder /var/task/bin/cli /var/task/bin/cli
COPY --from=builder /var/task/assets/ /var/task/assets/
ENTRYPOINT ["/var/task/bin/cli"]
