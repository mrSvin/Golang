FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o /docker-rest
EXPOSE 8090
CMD ["/docker-rest"]