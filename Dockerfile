# ----------------------------------------------------
# Base
# ----------------------------------------------------
FROM golang:1.16.3-alpine3.13 AS build_base
RUN mkdir -p {/app/bin}
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

# ----------------------------------------------------
# Build + Test
# ----------------------------------------------------
FROM build_base AS build
WORKDIR /app
COPY --from=build_base /app .
COPY . .

ENV GOOS linux
ENV GOARCH amd64
ENV CGO_ENABLED 0
RUN go build -ldflags "-s -w " -o app main.go

# ----------------------------------------------------
# Release
# ----------------------------------------------------
FROM alpine:3.13 AS release
ENV LC_ALL=en_US.UTF-8
ENV LC_LANG=en_US.UTF-8
ENV LC_LANGUAGE=en_US.UTF-8

RUN mkdir -p {/app/bin}
COPY --from=build /app/app /app/bin/app

WORKDIR /app
ENTRYPOINT ["/app/bin/app"]

EXPOSE 8080