FROM alpine:latest
RUN apk add libc6-compat
WORKDIR /app
COPY bin/pokebattle .
RUN chmod +x bin/pokebattle
CMD [ "/app/pokebattle" ]