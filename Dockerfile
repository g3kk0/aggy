FROM golang as builder

WORKDIR /build

COPY . .

RUN GCO_ENABLED=0 GOOS=linux go build -o aggy


FROM alpine:latest

RUN apk add --no-cache ca-certificates

COPY --from=builder /build/aggy /aggy
RUN chmod +x /aggy

ENV PORT 8080

CMD ["/aggy"]
