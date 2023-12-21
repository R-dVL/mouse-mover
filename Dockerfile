FROM golang:1.21

WORKDIR /usr/src/app

COPY . .

RUN go mod download && go mod verify

RUN go build -v -o /usr/local/bin/mouse_mover ./cmd/mouse_mover/main.go

CMD ["mouse_mover"]