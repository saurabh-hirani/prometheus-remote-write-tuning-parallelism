FROM golang:alpine

WORKDIR /build

RUN apk update && apk add git curl make build-base

RUN git clone https://github.com/prometheus-community/avalanche.git

RUN cd /build/avalanche && make all

WORKDIR /dist

RUN cp /build/avalanche/avalanche .

EXPOSE 9001

CMD ["/dist/avalanche"]
