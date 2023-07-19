FROM golang:1.20-alpine3.18

WORKDIR /usr/src/computer-club-DES

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod ./
RUN go mod download && go mod verify

COPY *.go ./
RUN go build -v -o /usr/local/bin/computer-club-DES ./...

ENTRYPOINT [ "/usr/local/bin/computer-club-DES" ]