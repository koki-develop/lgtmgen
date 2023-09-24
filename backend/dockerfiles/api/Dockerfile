FROM golang:1.21 AS base
WORKDIR /var/task

ENV CGO_CFLAGS_ALLOW="-Xpreprocessor"
ENV GOARCH=amd64
ENV GOOS=linux

RUN apt update

# Install ImageMagick
RUN wget https://imagemagick.org/archive/ImageMagick.tar.gz -O /tmp/ImageMagick.tar.gz \
  && tar -xzf /tmp/ImageMagick.tar.gz -C /tmp \
  && cd /tmp/ImageMagick-* \
  && ./configure \
  && make \
  && make install \
  && ldconfig /usr/local/lib

RUN apt install -y libjpeg-dev

# ----------------------------
FROM base AS builder

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download -x

# Build
COPY . .
RUN go build -ldflags="-s -w" -o /var/task/bin/cli .

# ----------------------------

FROM base

COPY --from=builder /var/task/bin/cli /var/task/bin/cli
ENTRYPOINT ["/var/task/bin/cli"]
