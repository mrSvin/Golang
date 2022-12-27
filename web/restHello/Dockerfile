FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o /docker-hello
EXPOSE 8080
CMD ["/docker-hello"]