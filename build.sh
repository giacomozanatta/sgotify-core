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