FROM ubuntu
COPY server /bin
WORKDIR /bin
EXPOSE 8080
CMD [ "server" ]