FROM golang:alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
RUN go mod download

COPY . .

RUN go build -o remote-write-receiver .

WORKDIR /dist

RUN cp /build/remote-write-receiver .

EXPOSE 4321

CMD ["/dist/remote-write-receiver"]
