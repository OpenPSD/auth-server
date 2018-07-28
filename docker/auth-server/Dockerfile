FROM alpine:3.7

ENV AUTH_HYDRA_URL="http://localhost:4444"  \
    AUTH_PORT="8000"

EXPOSE 8000
COPY auth-server /auth-server
COPY web /web
VOLUME ["/data", "/config"]
WORKDIR /

CMD ["./auth-server"]