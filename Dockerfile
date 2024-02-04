ARG OS
ARG ARCH

FROM ghcr.io/r-dvl/golang-builder:${OS}-${ARCH}

ENV PROJECT_NAME=stay-active
ENV VERSION=x.x.x

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

RUN mkdir -p ./bin/${PROJECT_NAME}-${VERSION}.${GOOS}-${GOARCH}

# Compile binaries
CMD ["sh", "-c", "go build -o ./bin/${PROJECT_NAME}-${VERSION}.${GOOS}-${GOARCH}/${PROJECT_NAME}${EXT} ./cmd/stay_active/"]