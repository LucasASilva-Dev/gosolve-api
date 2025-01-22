FROM golang:1.23.5-alpine3.21 AS build-env

ARG VERSION
ARG COMMIT

RUN apk add --no-cache git

RUN mkdir /go/src/gosolve
WORKDIR /go/src/gosolve

COPY . /go/src/gosolve

RUN CGO_ENABLED=0 go build -a -mod=vendor -installsuffix cgo -ldflags="-w -s -X gosolve/internal/app.Version=$VERSION -X gosolve/internal/app.Commit=$COMMIT" -o "./build/gosolve" ./cmd/gosolve

FROM alpine:latest

COPY --from=build-env /go/src/gosolve/data/ /app/
COPY --from=build-env /go/src/gosolve/build/gosolve /gosolve/gosolve

WORKDIR /gosolve

ENTRYPOINT ["/gosolve/gosolve", "server"]
