FROM alpine:latest
RUN apk add libc6-compat
WORKDIR /app
COPY bin/pokebattle .
CMD [ "/app/pokebattle" ]