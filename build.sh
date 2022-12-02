ls -a
mkdir bin
go build -o bin/sgotify
#RUN cd ..
mkdir bin/client
cp -R client/css bin/client/css
cp -R client/templates bin/client/templates
mkdir bin/client/scripts
#RUN cp -R src/client/templates app/client/templates
gopherjs build client/scripts/main.go
cp client/scripts/main.js bin/client/scripts/main.js
cp client/scripts/main.js.map bin/client/scripts/main.js.map
#RUN ls -a
#RUN cd client/scripts
ls -a
cd bin
ls -a