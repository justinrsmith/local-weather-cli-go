FROM scratch
COPY local-weather /bin/local-weather
ENTRYPOINT ["/bin/local-weather"]
