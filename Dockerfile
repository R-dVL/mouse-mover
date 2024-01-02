ARG OS

FROM ghcr.io/r-dvl/golang-builder:${OS}

ENV TAG=

WORKDIR /home/app

# Install Robotgo dependencies
RUN apt update -y && apt install -y \
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
        libxkbcommon-dev

COPY . .

RUN go mod download && go mod verify

# Compile binaries
CMD ["sh", "-c", "go build -o ./bin/stay_active-${TAG}.${GOOS}-${GOARCH}.${EXT} ./cmd/stay_active/"]