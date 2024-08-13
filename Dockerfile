# Build all services

FROM golang:1.22-alpine AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn GO111MODULE=on go mod download \
        && apk add --no-cache make git\
        && cd app/admin \
        && for p in case group judge problem user contest; do cd $p && make build && cd ..; done \
        && cd .. \
        && cd user/contest && make build && cd ../.. \
        && cd public/auth && make build && cd ../.. \
        && cd gojudge && make build && cd ..


FROM alpine:edge AS case

COPY --from=builder /src/app/admin/case/bin/case /
EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf
CMD ["/case", "-conf", "/data/conf"]


FROM alpine:edge AS group

COPY --from=builder /src/app/admin/group/bin/group /
EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf
CMD ["/group", "-conf", "/data/conf"]


FROM alpine:edge AS judge

COPY --from=builder /src/app/admin/judge/bin/judge /
EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf
CMD ["/judge", "-conf", "/data/conf"]


FROM alpine:edge AS problem

COPY --from=builder /src/app/admin/problem/bin/problem /
EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf
CMD ["/problem", "-conf", "/data/conf"]


FROM alpine:edge AS user

COPY --from=builder /src/app/admin/user/bin/user /
EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf
CMD ["/user", "-conf", "/data/conf"]


FROM alpine:edge AS contest

COPY --from=builder /src/app/admin/contest/bin/contest /
EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf
CMD ["/contest", "-conf", "/data/conf"]


FROM alpine:edge AS user-contest

COPY --from=builder /src/app/user/contest/bin/contest /
EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf
CMD ["/contest", "-conf", "/data/conf"]


FROM alpine:edge AS public-auth

COPY --from=builder /src/app/public/auth/bin/auth /
EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf
CMD ["/auth", "-conf", "/data/conf"]


FROM alpine:edge AS gojudge-server

COPY --from=builder /src/app/gojudge/bin/gojudge /
VOLUME /data/conf
CMD ["/gojudge", "-conf", "/data/conf"]