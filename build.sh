go install github.com/gopherjs/gopherjs@v1.18.0-beta1
go install golang.org/dl/go1.18.6@latest
go1.18.6 download
export GOPHERJS_GOROOT="$(go1.18.6 env GOROOT)"
ls -a
mkdir bin
go build -o bin/sgotify
#RUN cd ..
#RUN cp -R src/client/css app/client/css
#RUN cp -R src/client/templates app/client/templates
gopherjs build client/scripts/main.go
#RUN ls -a
#RUN cd client/scripts
ls -a
cd bin
ls -a