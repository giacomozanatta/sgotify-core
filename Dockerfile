FROM golang:latest
WORKDIR /app
COPY . ./

RUN ./build.sh
