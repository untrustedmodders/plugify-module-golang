cd test/example_plugin/
go clean -cache
go build -buildmode=plugin -tags=plugin -o example_plugin.so  ./
cd ../..
cd test/cross_call_worker/
go clean -cache
go build -buildmode=plugin -tags=plugin -o cross_call_worker.so  ./
cd ../..
cd managed
go clean -cache
go build -buildmode=c-shared -o libplugify-host-golang.so  ./
cd ..