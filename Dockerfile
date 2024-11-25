# Build all services

FROM golang:1.22-alpine AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn GO111MODULE=on go mod download \
        && apk add --no-cache make git\
        && cd app/admin/admin && make build && cd ../.. \
        && cd user/contest && make build && cd ../.. \
        && cd public/auth && make build && cd ../.. \
        && cd judge/gojudge && make build && cd ../.. \
        && cd judge/freshcup && make build && cd ../..


FROM alpine:edge AS admin

COPY --from=builder /src/app/admin/admin/bin/admin /
EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf
CMD ["/admin", "-conf", "/data/conf/config.yaml"]


FROM alpine:edge AS contest

COPY --from=builder /src/app/user/contest/bin/contest /
EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf
CMD ["/contest", "-conf", "/data/conf/config.yaml"]


FROM alpine:edge AS public-auth

COPY --from=builder /src/app/public/auth/bin/auth /
EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf
CMD ["/auth", "-conf", "/data/conf/config.yaml"]


FROM alpine:edge AS freshcup-server

COPY --from=builder /src/app/judge/freshcup/bin/freshcup /
VOLUME /data/conf
CMD ["/freshcup", "-conf", "/data/conf/config.yaml"]


FROM golang:alpine AS go-judge-build

WORKDIR /go/judge

RUN apk update && apk add git bash

RUN git clone https://ghp.ci/https://github.com/criyle/go-judge.git .

RUN git checkout v1.8.0

RUN go mod download -x

RUN go generate ./cmd/go-judge/version \
    && CGO_ENABLE=0 go build -v -tags grpcnotrace,nomsgpack -o go-judge ./cmd/go-judge


FROM alpine:3.17 AS gojudge-server

USER root

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk update && apk add --no-cache \
    wget \
    py3-pip \
    gcc \
    g++ \
    musl-dev \
    libtool \
    rust \
    ruby \
    go=1.19.9-r0 \
    #mono \
    ghc \
    openjdk17 \
    #php7 \
    #php7-cli \
    #php7-common \
    #php7-openssl \
    #php7-mbstring \
    #php7-json \
    #php7-phar \
    #php7-tokenizer \
    #php7-xml \
    #php7-curl \
    #php7-fileinfo \
    #fpc \
    webkit2gtk-4.1

COPY --from=go-judge-build /go/judge/go-judge /usr/bin/sandbox
RUN chmod +x /usr/bin/sandbox

COPY --from=builder /src/app/judge/gojudge/bin/gojudge /
VOLUME /data/conf

CMD sh -c "/usr/bin/sandbox -enable-grpc -grpc-addr 127.0.0.1:5051 & /gojudge -conf /data/conf/config.yaml"