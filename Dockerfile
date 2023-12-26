FROM ubuntu:latest

WORKDIR /home/app

# Build params
ENV OS=
ENV EXTENSION=
ENV ARCH=
ENV TAG=
ENV CGO_ENABLED=0
ENV CC=

RUN apt update -y && apt upgrade -y

# Update ca-certificates
RUN apt update -y && apt install -y ca-certificates && \
        update-ca-certificates

# Install dependencies
RUN apt update -y && apt install -y \
        golang-go \
        git \
        gcc \
        libc6-dev \
        libx11-dev \
        xorg-dev \
        libxtst-dev \
        xsel \
        xclip \
        libpng++-dev \
        xcb \
        libxcb-xkb-dev \
        x11-xkb-utils \
        libx11-xcb-dev \
        libxkbcommon-x11-dev \
        libxkbcommon-dev \
        mingw-w64

COPY . .

RUN go mod download && go mod verify

# Compile Windowsx64 binaries
CMD ["sh", "-c", "GOOS=$OS GOARCH=$ARCH CGO_ENABLED=$CGO_ENABLED CC=$CC go build -o ./bin/stay_active-$TAG.$OS-${ARCH}${EXTENSION} ./cmd/stay_active/"]