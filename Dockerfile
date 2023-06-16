# syntax=docker/dockerfile:1.4
FROM cgr.dev/chainguard/go:latest as build
LABEL maintainer="Will Norris <will@willnorris.com>"

WORKDIR /app
COPY . .

ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -v -o redirect .

FROM cgr.dev/chainguard/static:latest

COPY --from=build /app/redirect /app/redirect

ENTRYPOINT ["/app/redirect"]

EXPOSE 8080
