mkdir bin
go build -o bin/sgotify

mkdir bin/client
cp -R client/css bin/client/css
cp -R client/templates bin/client/templates
mkdir bin/client/scripts

gopherjs build client/scripts/main.go
cp client/scripts/main.js bin/client/scripts/main.js
cp client/scripts/main.js.map bin/client/scripts/main.js.map
