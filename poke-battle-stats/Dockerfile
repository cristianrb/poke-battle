FROM alpine:latest
RUN apk add libc6-compat
WORKDIR /app
COPY bin/pokebattlestats .
RUN chmod +x pokebattlestats
CMD [ "/app/pokebattlestats" ]