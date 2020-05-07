# -----------------------------------------------------------------------------
# Stage...
FROM golang:1.14.2-alpine AS build

ENV VERSION=0.0.1
WORKDIR /app

COPY go.mod go.sum main.go /app/
COPY internal /app/internal/

RUN apk --no-cache add git ;\
  CGO_ENABLED=0 go build -o /go/bin/picker \
  -ldflags="-w -s -X main.version=$VERSION" /app/main.go


# -----------------------------------------------------------------------------
# Image...
FROM alpine:3.8

COPY --from=build /go/bin/picker /bin/picker
COPY assets assets

EXPOSE 4500
ENTRYPOINT ["/bin/picker", "-d"]