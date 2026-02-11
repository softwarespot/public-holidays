# See URL: https://hub.docker.com/_/golang
# Use the Go image to build the binary only
FROM golang:1.26.0 AS builder
ENV CGO_ENABLED=0
ENV GOOS=linux
WORKDIR /go/src/public-holidays/
COPY . .

# Overwrite the development .env with the .env.production
COPY .env.production .env

RUN make

# See URL: https://hub.docker.com/_/alpine
# Use this image (~50MB) to run the "public-holidays", as the Go image contains too much bloat,
# which isn't needed for running the application in production and the image which can be uploaded
# to a public/private Docker register is then small
FROM alpine:3.20.3

COPY --from=builder /go/src/public-holidays/bin/* ./
COPY --from=builder /go/src/public-holidays/.env ./
EXPOSE 10000
CMD ["./public-holidays"]
