FROM alpine

ADD ./micro_linux_386 /opt/micro

CMD ["/opt/micro"]
