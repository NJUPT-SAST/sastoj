# Build base with dependencies
FROM golang:1.22-alpine AS base
WORKDIR /src
COPY go.mod go.sum ./
RUN GOPROXY=https://goproxy.cn GO111MODULE=on go mod download
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk add --no-cache make git

# Define build arguments
ARG GIT_COMMIT=unknown
ARG SERVICE=all

# Admin service builder
FROM base AS admin-builder
ARG SERVICE
COPY . /src
WORKDIR /src/app/admin/admin
RUN if [ "$SERVICE" = "all" ] || [ "$SERVICE" = "admin" ]; then \
      make build && \
      echo "${GIT_COMMIT}" > /git_commit.txt; \
    else \
      echo "Skipping admin build" && \
      mkdir -p /src/app/admin/admin/bin && \
      touch /git_commit.txt; \
    fi

# Contest service builder
FROM base AS contest-builder
ARG SERVICE
COPY . /src
WORKDIR /src/app/user/contest
RUN if [ "$SERVICE" = "all" ] || [ "$SERVICE" = "contest" ]; then \
      make build && \
      echo "${GIT_COMMIT}" > /git_commit.txt; \
    else \
      echo "Skipping contest build" && \
      mkdir -p /src/app/user/contest/bin && \
      touch /git_commit.txt; \
    fi

# Auth service builder
FROM base AS auth-builder
ARG SERVICE
COPY . /src
WORKDIR /src/app/public/auth
RUN if [ "$SERVICE" = "all" ] || [ "$SERVICE" = "auth" ]; then \
      make build && \
      echo "${GIT_COMMIT}" > /git_commit.txt; \
    else \
      echo "Skipping auth build" && \
      mkdir -p /src/app/public/auth/bin && \
      touch /git_commit.txt; \
    fi

# Gojudge service builder
FROM base AS gojudge-builder
ARG SERVICE
COPY . /src
WORKDIR /src/app/judge/gojudge
RUN if [ "$SERVICE" = "all" ] || [ "$SERVICE" = "gojudge" ]; then \
      make build && \
      echo "${GIT_COMMIT}" > /git_commit.txt; \
    else \
      echo "Skipping gojudge build" && \
      mkdir -p /src/app/judge/gojudge/bin && \
      touch /git_commit.txt; \
    fi

# Freshcup service builder
FROM base AS freshcup-builder
ARG SERVICE
COPY . /src
WORKDIR /src/app/judge/freshcup
RUN if [ "$SERVICE" = "all" ] || [ "$SERVICE" = "freshcup" ]; then \
      make build && \
      echo "${GIT_COMMIT}" > /git_commit.txt; \
    else \
      echo "Skipping freshcup build" && \
      mkdir -p /src/app/judge/freshcup/bin && \
      touch /git_commit.txt; \
    fi

# Admin service image
FROM alpine:edge AS admin
COPY --from=admin-builder /src/app/admin/admin/bin/admin /
COPY --from=admin-builder /git_commit.txt /git_commit.txt
EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf
CMD ["/admin", "-conf", "/data/conf/config.yaml"]

# Contest service image
FROM alpine:edge AS contest
COPY --from=contest-builder /src/app/user/contest/bin/contest /
COPY --from=contest-builder /git_commit.txt /git_commit.txt
EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf
CMD ["/contest", "-conf", "/data/conf/config.yaml"]

# Auth service image
FROM alpine:edge AS public-auth
COPY --from=auth-builder /src/app/public/auth/bin/auth /
COPY --from=auth-builder /git_commit.txt /git_commit.txt
EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf
CMD ["/auth", "-conf", "/data/conf/config.yaml"]

# Freshcup service image
FROM alpine:edge AS freshcup-server
COPY --from=freshcup-builder /src/app/judge/freshcup/bin/freshcup /
COPY --from=freshcup-builder /git_commit.txt /git_commit.txt
VOLUME /data/conf
CMD ["/freshcup", "-conf", "/data/conf/config.yaml"]

# Go-judge builder
FROM golang:alpine AS go-judge-build
ARG SERVICE
WORKDIR /go/judge
RUN if [ "$SERVICE" = "all" ] || [ "$SERVICE" = "gojudge" ]; then \
      apk update && apk add git bash && \
      git clone https://ghfast.top/https://github.com/criyle/go-judge.git . && \
      git checkout v1.8.0 && \
      go mod download -x && \
      go generate ./cmd/go-judge/version && \
      CGO_ENABLE=0 go build -v -tags grpcnotrace,nomsgpack -o go-judge ./cmd/go-judge; \
    else \
      echo "Skipping go-judge build" && \
      mkdir -p /go/judge && \
      touch /go/judge/go-judge; \
    fi

# Gojudge service image
FROM alpine:3.17 AS gojudge-server
USER root
ARG SERVICE
RUN if [ "$SERVICE" = "all" ] || [ "$SERVICE" = "gojudge" ]; then \
      sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
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
      webkit2gtk-4.1; \
    fi
COPY --from=go-judge-build /go/judge/go-judge /usr/bin/sandbox
RUN chmod +x /usr/bin/sandbox || true
COPY --from=gojudge-builder /src/app/judge/gojudge/bin/gojudge /
COPY --from=gojudge-builder /git_commit.txt /git_commit.txt
VOLUME /data/conf
CMD sh -c "/usr/bin/sandbox -enable-grpc -grpc-addr 127.0.0.1:5051 & /gojudge -conf /data/conf/config.yaml"