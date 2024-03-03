FROM golang:1.22.0-alpine
WORKDIR /project
COPY . .
RUN go mod tidy
EXPOSE 7000
ENTRYPOINT [ "go","run","." ]