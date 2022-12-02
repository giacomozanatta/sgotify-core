FROM golang:latest
RUN go install github.com/gopherjs/gopherjs@v1.18.0-beta1
RUN go install golang.org/dl/go1.18.6@latest
RUN go1.18.6 download
RUN export GOPHERJS_GOROOT="$(go1.18.6 env GOROOT)"
RUN mkdir bin
RUN mkdir bin/client
RUN mkdir bin/css
RUN mkdir bin/scripts
RUN mkdir bin/templates
RUN go build main.go -o bin/sgotify
RUN cp -R client/css bin/client/css
RUN cp -R client/templates bin/client/templates
RUN gopherjs build client/scripts/main.go
RUN ls -a
RUN cd client/scripts
RUN ls -a