FROM alpine

RUN apk add --update \
  curl \
  && rm -rf /var/cache/apk/*


HEALTHCHECK --interval=30s --timeout=30s --retries=3 \
  CMD curl -si localhost:8000/health | grep 'HTTP/1.1 200 OK' > /dev/null

ADD ./micro_linux_386 /opt/micro
EXPOSE 8000

CMD ["/opt/micro"]
