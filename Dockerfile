FROM gcr.io/distroless/static

USER nonroot:nonroot

COPY go-template /go-template

ENTRYPOINT ["/go-template"]
