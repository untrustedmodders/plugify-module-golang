cd plugify/
gcc -x c++ -c plugify.cc
ar rc libplugify.a plugify.o
cd ../
go build -buildmode=c-shared -o  awesome.dll main.go
pause