ARG BASE_IMAGE

FROM ${BASE_IMAGE} AS base
WORKDIR /var/task

ENV CGO_CFLAGS_ALLOW="-Xpreprocessor"
ENV GOARCH=amd64
ENV GOOS=linux

# Install dependencies
RUN apt update \
  && apt install -y \
  libjpeg-dev \
  libwebp-dev \
  ghostscript \
  build-essential \
  libfreetype6-dev

# Install ImageMagick
RUN wget https://imagemagick.org/archive/ImageMagick.tar.gz -O /tmp/ImageMagick.tar.gz \
  && tar -xzf /tmp/ImageMagick.tar.gz -C /tmp \
  && cd /tmp/ImageMagick-* \
  && ./configure \
  --with-freetype=yes \
  --with-webp=yes \
  && make \
  && make install \
  && ldconfig /usr/local/lib

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
COPY --from=builder /var/task/assets/ /var/task/assets/
ENTRYPOINT ["/var/task/bin/cli"]
