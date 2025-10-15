FROM golang:1.24-alpine3.20 as builder
WORKDIR /go/src/workspace
ADD . /go/src/workspace
RUN apk add --no-cache alpine-sdk && make all

# FROM gcr.io/distroless/static-debian12
FROM gcr.io/distroless/static-debian12:debug
WORKDIR /go/bin
COPY --from=builder /go/src/workspace/main .
USER 1001
CMD ["./main"]
