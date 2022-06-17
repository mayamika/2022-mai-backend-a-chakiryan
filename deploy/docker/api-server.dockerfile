ARG GO_VERSION=1.18.3-buster
ARG DEBIAN_VERSION=11.2-slim

FROM docker.io/golang:${GO_VERSION} as builder

COPY . /build
WORKDIR /build

RUN mkdir -p ./bin/
RUN go build -o ./bin/ /build/cmd/...

FROM docker.io/debian:${DEBIAN_VERSION}

COPY --from=builder /build/bin/api-server /

RUN apt-get update && apt-get install -y \
    wget

RUN mkdir --parents /certs && \
    wget "https://storage.yandexcloud.net/cloud-certs/CA.pem" \
    --output-document /certs/root.crt && \
    chmod 0600 /certs/root.crt

ENTRYPOINT [ "/api-server" ]
CMD []
