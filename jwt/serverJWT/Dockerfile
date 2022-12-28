FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o /docker-jwt
EXPOSE 8000
CMD ["/docker-jwt"]