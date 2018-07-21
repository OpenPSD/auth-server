FROM quay.io/coreos/dex:v2.10.0

COPY web /web
VOLUME ["/data", "/config"]
WORKDIR /

CMD ["serve", "/config/config.yaml"]