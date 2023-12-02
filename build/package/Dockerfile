FROM golang:1.19-alpine3.15 as build-env

WORKDIR /app

RUN apk add make git gcc g++

COPY . .

RUN /usr/bin/make build

FROM alpine:3.15

RUN apk add tzdata

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

WORKDIR /app

COPY configs/meta-egg-layout.yml /app/configs/meta-egg-layout.yml
COPY --from=build-env /app/build/bin/* /app/
CMD ["./meta-egg-layout"]
