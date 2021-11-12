FROM golang
MAINTAINER leiwingqueen
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o badminton-server src/main/server.go
CMD ["/app/badminton-server"]