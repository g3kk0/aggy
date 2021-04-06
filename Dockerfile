FROM golang:1.16.3-alpine as build-env

WORKDIR /workspace
COPY . /workspace
RUN go build -o aggy


FROM golang:1.16.3-alpine

WORKDIR /
COPY --from=build-env /workspace/aggy /
COPY --from=build-env /workspace/static /static

ENV PORT 8080

CMD ["/aggy"]
