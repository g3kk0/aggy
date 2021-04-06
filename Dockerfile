FROM golang as build-env

WORKDIR /workspace
ADD . /workspace
RUN GCO_ENABLED=0 go build -o aggy


FROM gcr.io/distroless/base-debian10

COPY --from=build-env /workspace/aggy /
ENV PORT 8080

CMD ["/aggy"]
