FROM golang:latest
RUN go install github.com/gopherjs/gopherjs@v1.18.0-beta1
RUN go install golang.org/dl/go1.18.6@latest
RUN go1.18.6 download
RUN export GOPHERJS_GOROOT="$(go1.18.6 env GOROOT)"
RUN mkdir app
RUN mkdir app/client
RUN mkdir app/css
RUN mkdir app/scripts
RUN mkdir app/templates
RUN ls -a
RUN go build src/main.go -o app/sgotify
RUN cp -R src/client/css app/client/css
RUN cp -R src/client/templates app/client/templates
RUN gopherjs build src/client/scripts/main.go
RUN ls -a
RUN cd client/scripts
RUN ls -a