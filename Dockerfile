FROM golang:latest

RUN mkdir /app
ADD . /goprog


WORKDIR /app
# User/local/bin/go




CMD ["app/main"]

EXPOSE 8080