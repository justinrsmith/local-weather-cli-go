FROM alpine
RUN apk add ca-certificates
COPY local-weather /bin/local-weather
ENTRYPOINT ["/bin/local-weather"]
