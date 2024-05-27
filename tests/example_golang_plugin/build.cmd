cd plugify/
gcc -x c++ -c plugify.cc -D_GLIBCXX_USE_CXX11_ABI=0
ar rc libplugify.a plugify.o
cd ../
go build -buildmode=c-shared -o  awesome.dll main.go
pause