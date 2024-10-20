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


FROM alpine:edge AS gojudge-server

COPY --from=builder /src/app/judge/gojudge/bin/gojudge /
VOLUME /data/conf
CMD ["/gojudge", "-conf", "/data/conf/config.yaml"]

FROM alpine:edge AS freshcup-server

COPY --from=builder /src/app/judge/freshcup/bin/freshcup /
VOLUME /data/conf
CMD ["/freshcup", "-conf", "/data/conf/config.yaml"]