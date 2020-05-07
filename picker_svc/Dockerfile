# -----------------------------------------------------------------------------
# Stage...
FROM golang:1.14.2-alpine AS build

<<!!YOUR_CODE!!>> -- Setup your version and build command

# -----------------------------------------------------------------------------
# Image...
FROM alpine:3.8

<<!!YOUR_CODE!!>> -- Build the exec from the previous stage in /bin

EXPOSE 4500
ENTRYPOINT ["/bin/picker", "-d"]