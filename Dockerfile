FROM gcr.io/distroless/cc
COPY tproxy /app/
CMD ["/app/tproxy"]
