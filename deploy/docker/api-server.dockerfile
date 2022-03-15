ARG GO_VERSION=1.18rc1-buster
ARG DEBIAN_VERSION=11.2-slim

FROM docker.io/golang:${GO_VERSION} as builder

COPY . /build
WORKDIR /build

RUN mkdir -p ./bin/
RUN go build -o ./bin/ /build/cmd/...

FROM docker.io/debian:${DEBIAN_VERSION}

COPY --from=builder /build/bin/api-server /

EXPOSE 8080
ENTRYPOINT [ "/api-server" ]
CMD [ "" ]
