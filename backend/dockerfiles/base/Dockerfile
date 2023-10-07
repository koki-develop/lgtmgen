FROM golang:1.21
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
