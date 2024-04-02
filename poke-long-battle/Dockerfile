FROM alpine:latest
RUN apk add libc6-compat
WORKDIR /app
COPY bin/pokelongbattle .
RUN chmod +x pokelongbattle
CMD [ "/app/pokelongbattle" ]