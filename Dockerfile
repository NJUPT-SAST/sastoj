FROM golang:1.22 AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn GO111MODULE=on go mod download \
        && cd app/admin \
        && for p in case group judge problem user contest; do cd $p && make build && cd ..; done


FROM debian:stable-slim AS case

COPY --from=builder /src/app/admin/case/bin/case /
EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf
CMD ["/case", "-conf", "/data/conf"]


FROM debian:stable-slim AS group

COPY --from=builder /src/app/admin/group/bin/group /
EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf
CMD ["/group", "-conf", "/data/conf"]


FROM debian:stable-slim AS judge

COPY --from=builder /src/app/admin/judge/bin/judge /
EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf
CMD ["/judge", "-conf", "/data/conf"]


FROM debian:stable-slim AS problem

COPY --from=builder /src/app/admin/problem/bin/problem /
EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf
CMD ["/problem", "-conf", "/data/conf"]


FROM debian:stable-slim AS user

COPY --from=builder /src/app/admin/user/bin/user /
EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf
CMD ["/user", "-conf", "/data/conf"]


FROM debian:stable-slim AS contest

COPY --from=builder /src/app/admin/contest/bin/contest /
EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf
CMD ["/contest", "-conf", "/data/conf"]