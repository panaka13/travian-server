# syntax=docker/dockerfile:1

FROM golang:1.17 as BUILD

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY db ./db
COPY handler ./handler
COPY model ./model
COPY server ./server
COPY Makefile .
RUN make


# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# WORKDIR /root/

# COPY --from=BUILD /app/build/main main
COPY cert ./cert
CMD ["build/main"]
