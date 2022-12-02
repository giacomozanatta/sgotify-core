FROM golang:latest
RUN go install github.com/gopherjs/gopherjs@v1.18.0-beta1 &&\
    go install golang.org/dl/go1.18.6@latest &&\
    go1.18.6 download &&\
    export GOPHERJS_GOROOT="$(go1.18.6 env GOROOT)"
WORKDIR /app
COPY . ./
CMD ./build.sh
