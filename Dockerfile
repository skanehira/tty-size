FROM busybox:1.30
COPY tty-size /usr/local/bin/tty-size
ENTRYPOINT ["tty-size"]
