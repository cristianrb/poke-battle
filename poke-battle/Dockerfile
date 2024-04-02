FROM alpine:latest
RUN apk add libc6-compat
WORKDIR /app
COPY bin/pokebattle .
RUN chmod +x pokebattle
CMD [ "/app/pokebattle" ]