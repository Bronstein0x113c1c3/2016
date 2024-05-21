FROM ubuntu
COPY /songslist /bin/songslist
COPY servers /bin

EXPOSE 8080/udp

WORKDIR /bin
CMD ["servers"]