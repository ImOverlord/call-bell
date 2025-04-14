FROM gcr.io/distroless/static

USER nonroot:nonroot

COPY call-bell /call-bell
COPY views /views

ENTRYPOINT ["/call-bell"]
