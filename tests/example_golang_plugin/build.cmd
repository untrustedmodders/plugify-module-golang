cd plugify/
gcc -c plugify.c
ar rc libplugify.a plugify.o
cd ../
go build -buildmode=c-shared -o awesome.dll -ldflags=-w main.go
pause